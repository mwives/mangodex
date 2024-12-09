package ui

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

func InsertRange(downloadType string) (float64, float64, error) {
	var downloadRange string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Description(fmt.Sprintf("Enter the range of %s you want to download", downloadType)),
			huh.NewInput().
				Value(&downloadRange).
				Validate(func(str string) error {
					// str must be N-M, N- or N
					pattern := `^(\d+-\d*|\d+)$`
					matched, _ := regexp.MatchString(pattern, str)
					if !matched {
						return fmt.Errorf("invalid range format")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		return 0, 0, err
	}

	start, end, err := convertStrRangeToFloat64(downloadRange)
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}

func convertStrRangeToFloat64(rangeStr string) (float64, float64, error) {
	var start, end float64
	var err error

	// if the range doesn't contain a dash, it's a single value
	if !strings.Contains(rangeStr, "-") {
		start, err = strconv.ParseFloat(rangeStr, 64)
		if err != nil {
			return 0, 0, err
		}
		return start, start, nil
	}

	ranges := strings.Split(rangeStr, "-")

	if ranges[1] == "" {
		ranges[1] = "0"
	}

	start, err = strconv.ParseFloat(ranges[0], 64)
	if err != nil {
		return 0, 0, err
	}

	end, err = strconv.ParseFloat(ranges[1], 64)
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}
