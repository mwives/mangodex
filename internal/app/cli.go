package app

import (
	"fmt"

	"github.com/mwives/mangodex/internal/ui"
	"github.com/mwives/mangodex/pkg/mangadex"
)

func Run() {
	mangadexApiClient := mangadex.NewMangadexApiClient("https://api.mangadex.org")

	searchType, searchQuery := ui.AskSearchType()

	var mangaList []mangadex.MangaResult
	var err error

	if searchType == ui.MangaTitleSearchType {
		mangaList, err = mangadexApiClient.SearchMangaByTitle(searchQuery)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if searchType == ui.MangaAuthorSearchType {
		authorResults, err := mangadexApiClient.SearchAuthors(searchQuery)
		if err != nil {
			fmt.Println(err)
			return
		}

		selectedAuthor, err := ui.SelectAuthor(authorResults)
		if err != nil {
			fmt.Println(err)
			return
		}

		mangaList, err = mangadexApiClient.SearchMangaByAuthorID(selectedAuthor.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Invalid search type")
		return
	}

	selectedManga, err := ui.SelectManga(mangaList)
	if err != nil {
		fmt.Println(err)
		return
	}

	ui.ConfirmMangaSelection(selectedManga)
	language, err := ui.SelectLanguage(*selectedManga.AvailableTranslatedLanguages)
	if err != nil {
		fmt.Println(err)
		return
	}
	conversionType, err := ui.SelectConversionType()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Downloading manga...")
	fmt.Println("Manga Title:", selectedManga.Title)
	fmt.Println("Language:", language)
	fmt.Println("Conversion Type:", conversionType)

}
