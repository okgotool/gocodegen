package app

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	RUNTIME_PATH = GetCurrentDirectory()
)

func GetRunningPath() string {
	return GetCurrentDirectory() + "/../"
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Failed to GetCurrentDirectory: " + err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}
