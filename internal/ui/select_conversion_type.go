package ui

import "github.com/charmbracelet/huh"

func SelectConversionType() (string, error) {
	var conversionType string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Conversion Type").
				Options(
					huh.NewOption("PDF", "pdf"),
					huh.NewOption("MOBI", "mobi"),
					huh.NewOption("EPUB", "epub"),
					huh.NewOption("ZIP", "zip"),
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
