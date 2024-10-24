package mangadex

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type MangadexApiClient struct {
	client *resty.Client
	apiURL string
}

func NewMangadexApiClient(baseURL string) *MangadexApiClient {
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
