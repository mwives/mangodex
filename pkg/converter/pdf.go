package converter

import (
	"fmt"
	"image"
	"os"

	"github.com/go-pdf/fpdf"
)

func createPDF(inputDir string, outputPath string) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	defer pdf.Close()

	files, err := os.ReadDir(inputDir)
	if err != nil {
		return err
	}

	sortPagesAndChapters(files)

	for _, file := range files {
		imgPath := fmt.Sprintf("%s/%s", inputDir, file.Name())
		img, err := os.Open(imgPath)
		if err != nil {
			return err
		}
		defer img.Close()

		imgConfig, _, err := image.DecodeConfig(img)
		if err != nil {
			return err
		}

		pdf.AddPageFormat("P", fpdf.SizeType{
			Wd: float64(imgConfig.Width),
			Ht: float64(imgConfig.Height),
		})
		pdf.ImageOptions(
			imgPath, 0, 0,
			float64(imgConfig.Width), float64(imgConfig.Height),
			false,
			fpdf.ImageOptions{
				ReadDpi: true,
			}, 0, "",
		)
	}

	return pdf.OutputFileAndClose(outputPath)
}
