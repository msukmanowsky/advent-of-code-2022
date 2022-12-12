package utils

import (
	"path"
	"runtime"
)

func ModRelativeFilePath(elem ...string) string {
	_, filename, _, _ := runtime.Caller(1)
	paths := append([]string{path.Dir(filename)}, elem...)
	filepath := path.Join(paths...)
	return filepath
}
