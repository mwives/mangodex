package mangadex

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
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

	_, err := m.client.R().
		SetQueryParam("limit", "10").
		SetQueryParam("offset", "0").
		SetQueryParam("title", title).
		SetQueryParam("includes[]", "author").
		SetResult(&result).
		Get("/manga")
	if err != nil {
		return nil, err
	}

	mangaResults := FormatMangaResult(result)
	return mangaResults, nil
}

func (m *MangadexApiClient) SearchMangaByAuthorID(authorID string) ([]MangaResult, error) {
	var result mangadexMangaSearchResult

	_, err := m.client.R().
		SetQueryParam("limit", "10").
		SetQueryParam("offset", "0").
		SetQueryParam("authorOrArtist", authorID).
		SetQueryParam("includes[]", "author").
		SetResult(&result).
		Get("/manga")
	if err != nil {
		return nil, err
	}

	mangaResults := FormatMangaResult(result)
	return mangaResults, nil
}

func (m *MangadexApiClient) SearchAuthors(name string) ([]Author, error) {
	var result mangadexAuthorSearchResult

	_, err := m.client.R().
		SetQueryParam("limit", "10").
		SetQueryParam("offset", "0").
		SetQueryParam("name", name).
		SetResult(&result).
		Get("/author")
	if err != nil {
		return nil, err
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

	_, err := m.client.R().
		SetResult(&result).
		SetQueryParam("translatedLanguage[]", language).
		Get(fmt.Sprintf("/manga/%s/aggregate", mangaID))
	if err != nil {
		return MangaAggregate{}, err
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

	_, err := m.client.R().
		SetResult(&result).
		Get(fmt.Sprintf("/at-home/server/%s?forcePort443=true", chapterID))
	if err != nil {
		return ChapterData{}, err
	}

	chapterData := ChapterData{
		Hash: result.Chapter.Hash,
		Data: result.Chapter.Data,
	}
	return chapterData, nil
}
