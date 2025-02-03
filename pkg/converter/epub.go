package converter

import (
	"fmt"
	"os"
	"os/exec"
)

func createEPUB(inputDir, outputPath string) error {
	if !verifyKCCInstalled() {
		fmt.Println("Error: KCC (Kindle Comic Converter) is not installed.")
		fmt.Println("Please install it by following the instructions at: https://github.com/ciromattia/kcc")
		os.Exit(1)
	}

	cmd := exec.Command(
		"kcc-c2e",
		"--manga-style",
		"--hq",
		"--upscale",
		"--splitter", "1",
		"--cropping", "0",
		inputDir,
		"--output", outputPath,
		"--title", "TODO TITLE",
		"--author", "TODO AUTHOR",
		"--format", "EPUB",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to convert to EPUB: %s", string(output))
	}

	fmt.Println("KCC Conversion Successful:", outputPath)
	return nil
}

func verifyKCCInstalled() bool {
	_, err := exec.LookPath("kcc-c2e")
	return err == nil
}
