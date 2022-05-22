package internal

import (
	"path/filepath"
	"syscall"
)

func normalizePath(path string) string {
	if filepath.VolumeName(path) == "" && filepath.HasPrefix(path, "\\") {
		var err error
		path, err = syscall.FullPath(path)
		if err != nil {
			panic(err)
		}
	}
	return path
}
