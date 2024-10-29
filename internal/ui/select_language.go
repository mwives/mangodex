package ui

import (
	"github.com/charmbracelet/huh"
)

func SelectLanguage(availableLanguages []string) (string, error) {
	var selectedLanguage string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Language").
				OptionsFunc(func() []huh.Option[string] {
					var options []huh.Option[string]
					for _, language := range availableLanguages {
						options = append(options, huh.NewOption(languageToFlag(language), language))
					}
					return options
				}, &availableLanguages).
				Value(&selectedLanguage),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return selectedLanguage, nil
}
