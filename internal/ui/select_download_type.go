package ui

import "github.com/charmbracelet/huh"

var (
	DownloadByVolumeType  = "volume"
	DownloadByChapterType = "chapter"
)

func SelectDownloadType() (string, error) {
	var conversionType string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Conversion Type").
				Options(
					huh.NewOption("Download by Volume", DownloadByVolumeType),
					huh.NewOption("Download by Chapter", DownloadByChapterType),
				).
				Value(&conversionType),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return conversionType, nil
}
