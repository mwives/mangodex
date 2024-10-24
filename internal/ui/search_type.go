package ui

import (
	"log"

	"github.com/charmbracelet/huh"
)

var (
	searchType  string
	searchQuery string
	mangaId     string
)

var (
	MangaTitleSearchType  = "manga_title"
	MangaAuthorSearchType = "author"
)

func AskSearchType() (string, string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("How would you like to search?").
				Options(
					huh.NewOption("Manga Title", MangaTitleSearchType),
					huh.NewOption("Author Name", MangaAuthorSearchType),
				).
				Value(&searchType),
			huh.NewInput().
				TitleFunc(func() string {
					if searchType == MangaTitleSearchType {
						return "Enter Manga Title"
					} else if searchType == MangaAuthorSearchType {
						return "Enter Author Name"
					}
					return ""
				}, &searchType).Value(&searchQuery),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return searchType, searchQuery
}
