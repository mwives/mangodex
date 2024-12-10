package converter

import (
	"archive/zip"
	"io"
	"os"
)

func createZIP(inputDir, outputPath string) error {
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)

	err = addFilesToZIP(zipWriter, inputDir)
	if err != nil {
		return err
	}

	err = zipWriter.Close()
	if err != nil {
		return err
	}

	return nil
}

func addFilesToZIP(zipWriter *zip.Writer, inputDir string) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := inputDir + "/" + file.Name()
		if !file.IsDir() {
			err = addFileToZIP(zipWriter, filePath, file.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func addFileToZIP(zipWriter *zip.Writer, filePath, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	header.Name = fileName

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return nil
}
