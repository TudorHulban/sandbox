package archiving

import "os"

func createFile(filePath string) (*os.File, func(), error) {
	targetFile, errCr := os.Create(filePath + zipExtension)
	if errCr != nil {
		return nil, nil, errCr
	}

	return targetFile,
		func() {
			targetFile.Close()
		},
		nil
}

// see:
// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func fileExists(filePath string) bool {
	if _, errPath := os.Stat(filePath); errPath != nil {
		if os.IsNotExist(errPath) {
			return false
		}
	}

	return true
}
