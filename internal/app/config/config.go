package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetDefault("debug", false)
	viper.SetDefault("save_dir", GetDefaultSaveDir())

	viper.BindEnv("debug", "MANGODEX_DEBUG")
	viper.BindEnv("save_dir", "MANGODEX_SAVE_DIR")
}

func GetSaveDir() string {
	dir := viper.GetString("save_dir")
	if dir == "" {
		panic("save_dir is not configured")
	}
	return dir
}

// GetDefaultSaveDir determines the platform-specific default save directory
func GetDefaultSaveDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("unable to detect user's home directory")
	}
	// Save to:
	// Windows: %USERPROFILE%\Documents\MangoDex
	// Linux/MacOS: ~/Documents/MangoDex
	return filepath.Join(homeDir, "Documents", "MangoDex")
}
