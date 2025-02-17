package main

import (
	"fmt"
	//"os"

	"onis-emem.com/go/files/fileutils"
)

func main() {
	// optionally we can get the root path with os.Getwd() but it works without it
	// rootPath, _ := os.Getwd()
	//c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	filePath := "data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)

	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}
}
