package converter

type Converter interface {
	ConvertToPDF(inputDir, outputPath string) error
	ConvertToEPUB(inputDir, outputPath string) error
	ConvertToZIP(inputDir, outputPath string) error
}

type MangaConverter struct{}

func NewMangaConverter() *MangaConverter {
	return &MangaConverter{}
}

func (m *MangaConverter) ConvertToPDF(inputDir, outputPath string) error {
	return createPDF(inputDir, outputPath)
}

func (m *MangaConverter) ConvertToEPUB(inputDir, outputPath string) error {
	// implementation
	return nil
}

func (m *MangaConverter) ConvertToZIP(inputDir, outputPath string) error {
	// implementation
	return nil
}
