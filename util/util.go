package util

import (
	"os"
	"path/filepath"
	"strings"
)

func IsTestMode() bool {
	wd, _ := os.Getwd()
	if strings.HasSuffix(wd, "test") {
		return true
	}
	return false
}

func RelPathToAbs(path string) string {
	if IsTestMode() {
		path = filepath.Join("..", path)
	}
	path, _ = filepath.Abs(path)
	return path
}
