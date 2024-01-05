package main

import (
	"bufio"
	"os"
	"strings"
)

type entry struct {
	folderInfo string
	indent     int
}

const _filePath = "folders.txt"

// readFile Helper reading file contents to a slice.
func readFile(filePath string) ([]string, error) {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}

func convertToEntry(line string) *entry {
	trimmed := strings.TrimLeft(line, " ")

	return &entry{
		folderInfo: trimmed,
		indent:     len(line) - len(trimmed),
	}
}

func convertToEntries(content []string) []*entry {
	result := make([]*entry, len(content))

	for ix, line := range content {
		result[ix] = convertToEntry(line)
	}

	return result
}

func parse(entries []*entry) []string {
	var res []string //nolint:prealloc

	var stackFolders stack

	var stackIndents stack

	for ix, entry := range entries {
		if ix == 0 {
			stackFolders.push(entry.folderInfo)
			stackIndents.push(0)

			res = append(res,
				commandMakeDirectory+space+stackFolders.String(),
			)

			continue
		}

		if entry.indent > stackIndents.peek().(int) {
			stackFolders.push(entry.folderInfo)
			stackIndents.push(entry.indent)

			res = append(res,
				commandMakeDirectory+space+stackFolders.String(),
			)

			continue
		}

		if entry.indent == stackIndents.peek().(int) {
			stackFolders.pop()

			stackFolders.push(entry.folderInfo)
			stackIndents.push(entry.indent)

			res = append(res,
				commandMakeDirectory+space+stackFolders.String(),
			)

			continue
		}

		for entry.indent < stackIndents.peek().(int) && len(stackIndents) > 1 {
			if entry.indent == stackIndents.peek().(int) {
				stackFolders.pop()

				break
			}

			stackFolders.pop()
			stackIndents.pop()
		}

		stackFolders.push(entry.folderInfo)
		stackIndents.push(entry.indent)

		res = append(res,
			commandMakeDirectory+space+stackFolders.String(),
		)
	}

	return res
}

func main() {
	// go test -run TestFile
}
