package mangadex

import "strconv"

func FormatMangaResult(mangaData mangadexMangaSearchResult) []MangaResult {
	var mangaResults []MangaResult
	for _, mangaData := range mangaData.Data {
		// Some manga titles are in Japanese, so we use the English title if available
		var mangaTitle = mangaData.Atributes.Title.En
		if mangaTitle == "" {
			mangaTitle = mangaData.Atributes.Title.Ja
		}

		// Concatenate all genres into a single string
		var genres string
		for _, tag := range mangaData.Atributes.Tags {
			if tag.Attributes.Group == "genre" {
				genres += tag.Attributes.Name.En + ", "
			}
		}
		if len(genres) > 0 {
			genres = genres[:len(genres)-2] // Remove the trailing comma and space
		}

		// Concatenate and add emoji flags to the alternative titles
		var altTitles string
		for _, altTitle := range mangaData.Atributes.AltTitles {
			if altTitle.En != "" {
				altTitles += "🇺🇸 " + altTitle.En + ", "
			}
			if altTitle.PtBr != "" {
				altTitles += "🇧🇷 " + altTitle.PtBr + ", "
			}
			if altTitle.JaRo != "" {
				altTitles += "🇯🇵 " + altTitle.JaRo + ", "
			}
		}
		if len(altTitles) > 0 {
			altTitles = altTitles[:len(altTitles)-2] // Remove the trailing comma and space
		}

		mangaResults = append(mangaResults, MangaResult{
			ID:                           mangaData.ID,
			Title:                        mangaTitle,
			Author:                       mangaData.Relationships[0].Attributes.Name,
			Genres:                       genres,
			AltTitles:                    altTitles,
			Year:                         mangaData.Atributes.Year,
			Status:                       mangaData.Atributes.Status,
			AvailableTranslatedLanguages: &mangaData.Atributes.AvailableTranslatedLanguages,
		})
	}

	return mangaResults
}

func FilterMangaVolumesAndChapters(mangaVolumesAndChapters MangaAggregate, startRange, endRange int) MangaAggregate {
	filteredMangaVolumesAndChapters := MangaAggregate{
		Volumes: []Volume{},
	}

	for _, volume := range mangaVolumesAndChapters.Volumes {
		volumeNumber, err := strconv.Atoi(volume.Volume)
		if err != nil {
			continue
		}

		if volumeNumber >= startRange && volumeNumber <= endRange {
			filteredMangaVolumesAndChapters.Volumes = append(filteredMangaVolumesAndChapters.Volumes, volume)
		}
	}

	return filteredMangaVolumesAndChapters
}
