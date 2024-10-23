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
	var result mangadexSearchResult

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
