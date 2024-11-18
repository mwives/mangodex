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

func FilterMangaVolumesAndChaptersByVolumeRange(mangaVolumesAndChapters MangaAggregate, startRange, endRange int) MangaAggregate {
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

	sortMangaVolumes(filteredMangaVolumesAndChapters)
	return filteredMangaVolumesAndChapters
}

func FilterMangaVolumesAndChaptersByChapterRange(mangaVolumesAndChapters MangaAggregate, startRange, endRange int, downloadType string) MangaAggregate {
	// Some chapters are not integers (ex: Berserk starts at chapter 0.1)
	startRangeFloat := float64(startRange)
	endRangeFloat := float64(endRange)

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

			if chapterNumber >= startRangeFloat && chapterNumber <= endRangeFloat {
				filteredChapters = append(filteredChapters, chapter)
			}
		}

		if len(filteredChapters) > 0 {
			filteredVolume := volume
			filteredVolume.Chapters = filteredChapters
			filteredMangaVolumesAndChapters.Volumes = append(filteredMangaVolumesAndChapters.Volumes, filteredVolume)
		}
	}

	sortMangaVolumes(filteredMangaVolumesAndChapters)
	return filteredMangaVolumesAndChapters
}

func sortMangaVolumes(mangaVolumesAndChapters MangaAggregate) {
	sort.Slice(mangaVolumesAndChapters.Volumes, func(i, j int) bool {
		volumeI, _ := strconv.Atoi(mangaVolumesAndChapters.Volumes[i].Volume)
		volumeJ, _ := strconv.Atoi(mangaVolumesAndChapters.Volumes[j].Volume)
		return volumeI < volumeJ
	})
	for _, volume := range mangaVolumesAndChapters.Volumes {
		sort.Slice(volume.Chapters, func(i, j int) bool {
			chapterI, _ := strconv.ParseFloat(volume.Chapters[i].Chapter, 64)
			chapterJ, _ := strconv.ParseFloat(volume.Chapters[j].Chapter, 64)
			return chapterI < chapterJ
		})
	}
}
