package helpers

import "os"

func GetFiles(fromFolder, withExtension string) ([]string, error) {
	entries, err := os.ReadDir(fromFolder)
	if err != nil {
		return []string{}, err
	}

	result := make([]string, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := fromFolder + "/" + entry.Name()

		result = append(result, path)
	}

	return result, nil
}
