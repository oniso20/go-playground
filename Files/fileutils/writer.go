package fileutils

import (
	"log"
	"os"
)

func WriteToFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
