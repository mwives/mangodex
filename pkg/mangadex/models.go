package mangadex

// Manga
type mangadexMangaSearchResult struct {
	Data []struct {
		ID        string `json:"id"`
		Atributes struct {
			Title struct {
				En string `json:"en"`
				Ja string `json:"ja"`
			} `json:"title"`
			AltTitles []struct {
				PtBr string `json:"pt-br"`
				En   string `json:"en"`
				JaRo string `json:"ja-ro"`
			} `json:"altTitles"`
			Description struct {
				En string `json:"en"`
			} `json:"description"`
			LastVolume             string `json:"lastVolume"`
			LastChapter            string `json:"lastChapter"`
			PublicationDemographic string `json:"publicationDemographic"`
			Status                 string `json:"status"`
			Year                   int    `json:"year"`
			Tags                   []struct {
				Attributes struct {
					Group string `json:"group"`
					Name  struct {
						En string `json:"en"`
					}
				}
			} `json:"tags"`
			AvailableTranslatedLanguages []string
		} `json:"attributes"`
		Relationships []struct {
			Attributes struct {
				Name string `json:"name"`
			} `json:"attributes"`
		} `json:"relationships"`
	}
}

type MangaResult struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genres    string `json:"genres"`
	AltTitles string `json:"altTitles"`
	Year      int    `json:"year"`
	Status    string `json:"status"`
}

// Author
type mangadexAuthorSearchResult struct {
	Data []struct {
		ID        string `json:"id"`
		Atributes struct {
			Name string `json:"name"`
		} `json:"attributes"`
	}
}

type Author struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TitleCount int    `json:"titleCount"`
}
