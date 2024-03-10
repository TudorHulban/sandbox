package archiving

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// see also https://gosamples.dev/zip-file/

type Zip struct{}

func (z *Zip) CompressFile(filePath string) error {
	fileTarget, closer, errCrFile := createFile(filePath)
	if errCrFile != nil {
		return errCrFile
	}
	defer closer()

	zipWriter := zip.NewWriter(fileTarget)
	defer zipWriter.Close()

	stats, errStat := fileTarget.Stat()
	if errStat != nil {
		return fmt.Errorf("stat: %w", errStat)
	}

	header, errHeader := zip.FileInfoHeader(stats)
	if errHeader != nil {
		return fmt.Errorf("header: %w", errHeader)
	}

	header.Name = filepath.Base(filePath)
	header.Method = zip.Deflate

	fileWriter, errCreateHeader := zipWriter.CreateHeader(header)
	if errCreateHeader != nil {
		return errCreateHeader
	}

	fileSource, errOpen := os.Open(filePath)
	if errOpen != nil {
		return errOpen
	}
	defer fileSource.Close()

	_, errCopy := io.Copy(fileWriter, fileSource)
	return errCopy
}
