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
