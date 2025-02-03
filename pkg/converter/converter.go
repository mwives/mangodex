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
		return createEPUB(inputDir, outputPath)
	case ZIPConversionType:
		return createZIP(inputDir, outputPath)
	default:
		return fmt.Errorf("unsupported conversion type: %s", conversionType)
	}
}
