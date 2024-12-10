package converter

import (
	"io/fs"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func sortPagesAndChapters(files []fs.DirEntry) {
	re := regexp.MustCompile(`\D`) // Match non-digit characters

	slices.SortFunc(files, func(a, b fs.DirEntry) int {
		// File names are in the format "chapter-page-uuid.jpg"
		// Extract file parts
		aParts := strings.Split(a.Name(), "-")
		bParts := strings.Split(b.Name(), "-")

		chapterA, errA := strconv.ParseFloat(aParts[0], 64)
		chapterB, errB := strconv.ParseFloat(bParts[0], 64)
		if errA != nil || errB != nil {
			return 0 // Skip invalid entries
		}

		pageA, errA := strconv.ParseFloat(re.ReplaceAllString(aParts[1], ""), 64)
		pageB, errB := strconv.ParseFloat(re.ReplaceAllString(bParts[1], ""), 64)
		if errA != nil || errB != nil {
			return 0 // Skip invalid entries
		}

		// Compare chapters
		if chapterA != chapterB {
			if chapterA < chapterB {
				return -1 // Sort before b
			}
			return 1 // Sort after b
		}

		// Compare pages
		if pageA != pageB {
			if pageA < pageB {
				return -1 // Sort before b
			}
			return 1 // Sort after b
		}

		// Equal
		return 0
	})
}
