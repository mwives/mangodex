package app

import (
	"fmt"

	"github.com/mwives/mangodex/internal/ui"
	"github.com/mwives/mangodex/pkg/mangadex"
)

func Run() {
	mangadexApiClient := mangadex.NewMangadexApiClient("https://api.mangadex.org")

	searchType, searchQuery := ui.RunSearch()

	var result []mangadex.MangaResult
	var err error

	if searchType == ui.MangaTitleSearchType {
		result, err = mangadexApiClient.SearchMangaByTitle(searchQuery)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if searchType == ui.MangaAuthorSearchType {
		// TODO: Implement search by author
		fmt.Println("Search by author not implemented yet")
	} else {
		fmt.Println("Invalid search type")
		return
	}

	fmt.Println(result)

}
