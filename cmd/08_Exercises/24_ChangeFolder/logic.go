package main

import "strings"

func handlePathLogic(path string) string {
	if len(path) == 0 {
		return ""
	}

	if path[0] != '/' {
		return currentFolder + "/" + path
	}

	elems := strings.Split(strings.Trim(path, " "), "/")

	result := make([]string, 0)

	for _, elem := range elems {
		if len(elem) == 0 {
			continue
		}

		if elem == ".." {
			result = result[:len(result)-1]

			if len(result) == 0 {
				result = []string{"/"}
			}

			if len(strings.Trim(result[0], " ")) == 0 {
				result = []string{"/"}
			}

			continue
		}

		result = append(result, elem)
	}

	if result[0] == "/" {
		return "/" + strings.Join(result[1:], "/")
	}

	return "/" + strings.Join(result, "/")
}
