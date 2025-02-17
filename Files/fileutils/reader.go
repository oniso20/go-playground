package fileutils

import "os"

func ReadTextFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		// Can't read file
		return "", err
	} else {
		// Read file
		return string(content), nil
	}
}
