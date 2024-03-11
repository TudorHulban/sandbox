package helpers

import (
	"os"
)

// see:
// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func FileExists(filePath string) bool {
	if _, errPath := os.Stat(filePath); errPath != nil {
		if os.IsNotExist(errPath) {
			return false
		}
	}

	return true
}

func DeleteFile(filePath string) error {
	if FileExists(filePath) {
		return nil
	}

	return os.Remove(filePath)
}

func CreateFile(filePath string, content string) error {
	return os.WriteFile(
		filePath,
		[]byte(content),
		0644,
	)
}
