package downloader

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/mwives/mangodex/internal/app/config"
	"github.com/mwives/mangodex/pkg/mangadex"
)

type Downloader struct {
	Client        *mangadex.MangadexApiClient
	UploadsClient *mangadex.MangadexUploadsApiClient
}

func NewDownloader(
	client *mangadex.MangadexApiClient,
	uploadsClient *mangadex.MangadexUploadsApiClient,
) *Downloader {
	return &Downloader{
		Client:        client,
		UploadsClient: uploadsClient,
	}
}

func (d *Downloader) DownloadChapter(chapter mangadex.Chapter) error {
	chapterData, err := d.Client.GetMangaChapterData(chapter.ID)
	if err != nil {
		return err
	}

	saveDir := config.GetSaveDir()
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create save directory: %w", err)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(chapterData.Data))

	// Each `data` element is a page file name (e.g. `1-[uuid].jpg`)
	for _, page := range chapterData.Data {
		wg.Add(1)

		go func(page string) {
			defer wg.Done()

			image, err := d.UploadsClient.FetchPageImage(chapterData.Hash, page)
			if err != nil {
				errChan <- fmt.Errorf("unable to fetch image: %w", err)
				return
			}

			if err := d.downloadFile(
				image, saveDir, fmt.Sprintf("%s-%s", chapter.Chapter, page),
			); err != nil {
				errChan <- fmt.Errorf("failed to download file for page %s: %w", page, err)
			}
		}(page)
	}

	wg.Wait()
	close(errChan)

	var downloadErrors []error
	for err := range errChan {
		downloadErrors = append(downloadErrors, err)
	}

	if len(downloadErrors) > 0 {
		return fmt.Errorf("failed to download chapter %s: %v", chapter.ID, downloadErrors)
	}

	return nil
}

func (d *Downloader) downloadFile(image []byte, saveDir, fileName string) error {
	filePath := filepath.Join(saveDir, fileName)
	fmt.Printf("Saving file to: %s\n", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(image))
	if err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", filePath, err)
	}

	return nil
}
