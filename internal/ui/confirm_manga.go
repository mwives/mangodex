package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/mwives/mangodex/pkg/mangadex"
)

func ConfirmMangaSelection(manga mangadex.MangaResult) (bool, error) {
	var confirmed bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Description(fmt.Sprintf(
					"Title: %s\nAuthor: %s\nGenres: %s\nAlt Titles: %s\nYear: %d\nStatus: %s",
					manga.Title,
					manga.Author,
					manga.Genres,
					manga.AltTitles,
					manga.Year,
					manga.Status,
				)),
			huh.NewConfirm().
				Title("Confirm selection?").
				Affirmative("Yes").
				Negative("No").
				Value(&confirmed),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println(err)
	}

	return confirmed, err
}
