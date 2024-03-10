package main

import (
	"os"
	"path/filepath"
)

func doPathWalk(folder string, files *[]string) error {
	fullPath, err := filepath.Abs(folder)
	if err != nil {
		return err
	}

	callback := func(path string, fi os.FileInfo, err error) error {
		return hashFile(fullPath, path, fi, files)
	}

	return filepath.Walk(fullPath, callback)
}

func hashFile(root, path string, fileInfo os.FileInfo, files *[]string) error {
	if fileInfo.IsDir() {
		return nil
	}

	if _, err := filepath.Rel(root, path); err != nil {
		return err
	}

	*files = append(*files, path)

	return nil
}
