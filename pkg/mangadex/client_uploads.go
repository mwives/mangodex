package mangadex

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type MangadexUploadsApiClient struct {
	client *resty.Client
	apiURL string
}

func NewMangadexUploadsApiClient() *MangadexUploadsApiClient {
	baseURL := "https://uploads.mangadex.org"

	client := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(10*time.Second).
		SetHeader("Content-Type", "application/json")

	return &MangadexUploadsApiClient{
		client: client,
		apiURL: baseURL,
	}
}

func (m *MangadexUploadsApiClient) FetchPageImage(chapterHash, pageFile string) ([]byte, error) {
	url := fmt.Sprintf("/data/%s/%s", chapterHash, pageFile)

	resp, err := m.client.R().
		SetHeader("Content-Type", "image/jpeg").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code %d for URL %s", resp.StatusCode(), url)
	}

	return resp.Body(), nil
}
