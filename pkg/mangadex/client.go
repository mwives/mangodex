package mangadex

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	mangaEndpoint            = "/manga"
	authorEndpoint           = "/author"
	chapterAggregateEndpoint = "/manga/%s/aggregate"
	chapterDataEndpoint      = "/at-home/server/%s"
)

type MangadexApiClient struct {
	client *resty.Client
	apiURL string
}

func NewMangadexApiClient() *MangadexApiClient {
	baseURL := "https://api.mangadex.org"

	client := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(10*time.Second).
		SetHeader("Content-Type", "application/json")

	return &MangadexApiClient{
		client: client,
		apiURL: baseURL,
	}
}

func (m *MangadexApiClient) SearchMangaByTitle(title string) ([]MangaResult, error) {
	var result mangadexMangaSearchResult

	resp, err := m.client.R().
		SetQueryParam("limit", "10").
		SetQueryParam("offset", "0").
		SetQueryParam("title", title).
		SetQueryParam("includes[]", "author").
		SetResult(&result).
		Get(mangaEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch manga by title: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	mangaResults := FormatMangaResult(result)
	return mangaResults, nil
}

func (m *MangadexApiClient) SearchMangaByAuthorID(authorID string) ([]MangaResult, error) {
	var result mangadexMangaSearchResult

	resp, err := m.client.R().
		SetQueryParam("limit", "10").
		SetQueryParam("offset", "0").
		SetQueryParam("authorOrArtist", authorID).
		SetQueryParam("includes[]", "author").
		SetResult(&result).
		Get(mangaEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch manga by author ID: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	mangaResults := FormatMangaResult(result)
	return mangaResults, nil
}

func (m *MangadexApiClient) SearchAuthors(name string) ([]Author, error) {
	var result mangadexAuthorSearchResult

	resp, err := m.client.R().
		SetQueryParam("limit", "10").
		SetQueryParam("offset", "0").
		SetQueryParam("name", name).
		SetResult(&result).
		Get(authorEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch authors by name: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var authorResults []Author
	for _, authorData := range result.Data {
		authorResults = append(authorResults, Author{
			ID:   authorData.ID,
			Name: authorData.Atributes.Name,
		})
	}

	return authorResults, nil
}

func (m *MangadexApiClient) SearchMangaVolumesAndChapters(mangaID, language string) (MangaAggregate, error) {
	var result mangadexMangaAggregateResult

	resp, err := m.client.R().
		SetResult(&result).
		SetQueryParam("translatedLanguage[]", language).
		Get(fmt.Sprintf(chapterAggregateEndpoint, mangaID))
	if err != nil {
		return MangaAggregate{}, fmt.Errorf("failed to fetch manga volumes and chapters: %w", err)
	}

	if resp.StatusCode() != 200 {
		return MangaAggregate{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var mangaAggregate MangaAggregate
	for volume, volumeData := range result.Volumes {
		var chapters []Chapter
		for chapter, chapterData := range volumeData.Chapters {
			chapters = append(chapters, Chapter{
				Chapter: chapter,
				ID:      chapterData.ID,
			})
		}

		mangaAggregate.Volumes = append(mangaAggregate.Volumes, Volume{
			Volume:   volume,
			Chapters: chapters,
		})
	}

	return mangaAggregate, nil
}

func (m *MangadexApiClient) GetMangaChapterData(chapterID string) (ChapterData, error) {
	var result mangadexChapterDataResult

	resp, err := m.client.R().
		SetResult(&result).
		Get(fmt.Sprintf(chapterDataEndpoint, chapterID))
	if err != nil {
		return ChapterData{}, fmt.Errorf("failed to fetch chapter data: %w", err)
	}

	if resp.StatusCode() != 200 {
		return ChapterData{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	chapterData := ChapterData{
		Hash: result.Chapter.Hash,
		Data: result.Chapter.Data,
	}
	return chapterData, nil
}
