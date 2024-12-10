package app

import (
	"fmt"
	"os"

	"github.com/mwives/mangodex/internal/app/config"
	"github.com/mwives/mangodex/internal/ui"
	"github.com/mwives/mangodex/pkg/converter"
	"github.com/mwives/mangodex/pkg/mangadex"
	"github.com/mwives/mangodex/pkg/mangadex/downloader"
)

func handleError(err error, userMessage string) {
	if os.Getenv("DEBUG") == "true" && err != nil {
		fmt.Println("DEBUG:", err)
	}
	fmt.Println(userMessage)
}

func Run() {
	mangadexApiClient := mangadex.NewMangadexApiClient()
	mangadexUploadsApiClient := mangadex.NewMangadexUploadsApiClient()
	downloader := downloader.NewDownloader(mangadexApiClient, mangadexUploadsApiClient)

	searchType, searchQuery := ui.SelectSearchType()

	var mangaList []mangadex.MangaResult
	var err error

	if searchType == ui.MangaTitleSearchType {
		mangaList, err = mangadexApiClient.SearchMangaByTitle(searchQuery)
		if err != nil {
			handleError(err, "There was an error during the title search. Please try again.")
			return
		}
	} else if searchType == ui.MangaAuthorSearchType {
		authorResults, err := mangadexApiClient.SearchAuthors(searchQuery)
		if err != nil {
			handleError(err, "There was an error during the author search. Please try again.")
			return
		}

		selectedAuthor, err := ui.SelectAuthor(authorResults)
		if err != nil {
			handleError(err, "There was an error during author selection. Please try again.")
			return
		}

		mangaList, err = mangadexApiClient.SearchMangaByAuthorID(selectedAuthor.ID)
		if err != nil {
			handleError(err, "There was an error searching for manga by author. Please try again.")
			return
		}
	} else {
		fmt.Println("Invalid search type")
		return
	}

	selectedManga, err := ui.SelectManga(mangaList)
	if err != nil {
		handleError(err, "There was an error during manga selection. Please try again.")
		return
	}

	mangaConfirmed, err := ui.ConfirmMangaSelection(selectedManga)
	if err != nil {
		handleError(err, "There was an error during manga confirmation. Please try again.")
		return
	}
	if !mangaConfirmed {
		fmt.Println("Manga selection was not confirmed. Exiting...")
		return
	}

	language, err := ui.SelectLanguage(*selectedManga.AvailableTranslatedLanguages)
	if err != nil {
		handleError(err, "There was an error during language selection. Please try again.")
		return
	}

	conversionType, err := ui.SelectConversionType()
	if err != nil {
		handleError(err, "There was an error during conversion type selection. Please try again.")
		return
	}

	downloadType, err := ui.SelectDownloadType()
	if err != nil {
		handleError(err, "There was an error during download type selection. Please try again.")
		return
	}

	// Here we have a problem. A manga may not have a volume or chapter for a certain language.
	// In this case, the best option would be to list all volumes/chapters that are available for the selected language.
	// For now, we will just assume that the manga always has volumes and chapters for the selected language.
	startRange, endRange, err := ui.InsertRange(downloadType)
	if err != nil {
		handleError(err, "There was an error specifying the download range. Please try again.")
		return
	}

	mangaVolumesAndChapters, err := mangadexApiClient.SearchMangaVolumesAndChapters(selectedManga.ID, language)
	if err != nil {
		handleError(err, "There was an error searching for manga volumes and chapters. Please try again.")
		return
	}

	rootSaveDir := config.GetSaveDir()

	if downloadType == ui.VolumeDownloadType {
		manga := mangadex.
			FilterMangaVolumesByRange(mangaVolumesAndChapters, startRange, endRange)

		for _, volume := range manga.Volumes {
			saveDir := fmt.Sprintf("%s/%s/%s", rootSaveDir, selectedManga.Title, fmt.Sprintf("Vol. %s", volume.Volume))
			for _, chapter := range volume.Chapters {
				downloader.DownloadChapter(chapter, saveDir)
			}
			converter.Convert(
				converter.ConversionType(conversionType),
				saveDir,
				fmt.Sprintf("%s/%s/%s.pdf", rootSaveDir, selectedManga.Title, fmt.Sprintf("%s - vol. %s", selectedManga.Title, volume.Volume)),
			)
			os.RemoveAll(saveDir)
		}
	} else {
		chapters := mangadex.
			FilterMangaChaptersByRange(mangaVolumesAndChapters, startRange, endRange)

		for _, chapter := range chapters {
			saveDir := fmt.Sprintf("%s/%s/%s", rootSaveDir, selectedManga.Title, fmt.Sprintf("Ch. %s", chapter.Chapter))
			downloader.DownloadChapter(chapter, saveDir)
			converter.Convert(
				converter.ConversionType(conversionType),
				saveDir,
				fmt.Sprintf("%s/%s/%s.pdf", rootSaveDir, selectedManga.Title, fmt.Sprintf("Ch. %s", chapter.Chapter)),
			)
			os.RemoveAll(saveDir)
		}
	}
}
