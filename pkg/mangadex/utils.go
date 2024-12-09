package mangadex

import (
	"sort"
	"strconv"
)

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

func FilterMangaVolumesByRange(mangaVolumesAndChapters MangaAggregate, startRange, endRange float64) MangaAggregate {
	filteredMangaVolumesAndChapters := MangaAggregate{
		Volumes: []Volume{},
	}

	for _, volume := range mangaVolumesAndChapters.Volumes {
		volumeNumber, err := strconv.ParseFloat(volume.Volume, 64)
		if err != nil {
			continue
		}

		if volumeNumber >= startRange && volumeNumber <= endRange {
			filteredMangaVolumesAndChapters.Volumes = append(filteredMangaVolumesAndChapters.Volumes, volume)
		}
	}

	sortMangaVolumes(&filteredMangaVolumesAndChapters)
	return filteredMangaVolumesAndChapters
}

func FilterMangaChaptersByRange(mangaVolumesAndChapters MangaAggregate, startRange, endRange float64) []Chapter {
	filteredMangaVolumesAndChapters := MangaAggregate{
		Volumes: []Volume{},
	}

	for _, volume := range mangaVolumesAndChapters.Volumes {
		filteredChapters := []Chapter{}
		for _, chapter := range volume.Chapters {
			chapterNumber, err := strconv.ParseFloat(chapter.Chapter, 64)
			if err != nil {
				continue
			}

			if chapterNumber >= startRange && chapterNumber <= endRange {
				filteredChapters = append(filteredChapters, chapter)
			}
		}

		if len(filteredChapters) > 0 {
			filteredVolume := volume
			filteredVolume.Chapters = filteredChapters
			filteredMangaVolumesAndChapters.Volumes = append(filteredMangaVolumesAndChapters.Volumes, filteredVolume)
		}
	}

	sortMangaVolumes(&filteredMangaVolumesAndChapters)
	return getChaptersFromManga(filteredMangaVolumesAndChapters)
}

func sortMangaVolumes(manga *MangaAggregate) {
	sort.Slice(manga.Volumes, func(i, j int) bool {
		volumeI, _ := strconv.Atoi(manga.Volumes[i].Volume)
		volumeJ, _ := strconv.Atoi(manga.Volumes[j].Volume)
		return volumeI < volumeJ
	})
	for _, volume := range manga.Volumes {
		sort.Slice(volume.Chapters, func(i, j int) bool {
			chapterI, _ := strconv.ParseFloat(volume.Chapters[i].Chapter, 64)
			chapterJ, _ := strconv.ParseFloat(volume.Chapters[j].Chapter, 64)
			return chapterI < chapterJ
		})
	}
}

func getChaptersFromManga(manga MangaAggregate) []Chapter {
	var chapters []Chapter
	for _, volume := range manga.Volumes {
		for _, chapter := range volume.Chapters {
			chapters = append(chapters, chapter)
		}
	}
	return chapters
}
