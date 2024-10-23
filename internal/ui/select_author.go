package ui

import (
	"github.com/charmbracelet/huh"
	"github.com/mwives/mangodex/pkg/mangadex"
)

func SelectAuthor(authorList []mangadex.Author) (mangadex.Author, error) {
	var selectedAuthor mangadex.Author

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[mangadex.Author]().
				Title("Select Author").
				OptionsFunc(func() []huh.Option[mangadex.Author] {
					var options []huh.Option[mangadex.Author]
					for _, author := range authorList {
						options = append(options, huh.NewOption(author.Name, author))
					}
					return options
				}, &authorList).
				Value(&selectedAuthor),
		),
	)

	err := form.Run()

	if err != nil {
		return mangadex.Author{}, err
	}

	return selectedAuthor, nil
}
