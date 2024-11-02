package ui

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

func InsertRange(downloadType string) (int, int, error) {
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

	start, end, err := convertStrRangeToInt(downloadRange)
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}

func convertStrRangeToInt(rangeStr string) (int, int, error) {
	var start, end int
	var err error

	// if the range doesn't contain a dash, it's a single value
	if !strings.Contains(rangeStr, "-") {
		start, err = strconv.Atoi(rangeStr)
		if err != nil {
			return 0, 0, err
		}
		return start, start, nil
	}

	ranges := strings.Split(rangeStr, "-")

	if ranges[1] == "" {
		ranges[1] = "0"
	}

	start, err = strconv.Atoi(ranges[0])
	if err != nil {
		return 0, 0, err
	}

	end, err = strconv.Atoi(ranges[1])
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}
