package converter

import "fmt"

type ConversionType string

const (
	PDFConversionType  ConversionType = "pdf"
	EPUBConversionType ConversionType = "epub"
	ZIPConversionType  ConversionType = "zip"
)

func Convert(conversionType ConversionType, inputDir, outputPath string) error {
	switch conversionType {
	case PDFConversionType:
		return createPDF(inputDir, outputPath)
	case EPUBConversionType:
		// not implemented yet
		return nil
	case ZIPConversionType:
		// not implemented yet
		return nil
	default:
		return fmt.Errorf("unsupported conversion type: %s", conversionType)
	}
}
