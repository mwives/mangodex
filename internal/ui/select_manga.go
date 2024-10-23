package ui

import (
	"github.com/charmbracelet/huh"
	"github.com/mwives/mangodex/pkg/mangadex"
)

func SelectManga(mangaList []mangadex.MangaResult) (mangadex.MangaResult, error) {
	var selectedManga mangadex.MangaResult

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[mangadex.MangaResult]().
				Title("Select Manga").
				OptionsFunc(func() []huh.Option[mangadex.MangaResult] {
					var options []huh.Option[mangadex.MangaResult]
					for _, manga := range mangaList {
						options = append(options, huh.NewOption(manga.Title, manga))
					}
					return options
				}, &mangaList).
				Value(&selectedManga),
		),
	)

	err := form.Run()

	if err != nil {
		return mangadex.MangaResult{}, err
	}

	return selectedManga, nil
}
