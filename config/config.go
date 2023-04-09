package config

import (
	"os"
	"path"
)

var mclConfigDir string

func init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return
	}
	mclConfigDir = path.Join(configDir, "MatsuriContestLog")
	os.MkdirAll(mclConfigDir, 0755)
}

func GetConfigFilePath(file string) string {
	return path.Join(mclConfigDir, file)
}
