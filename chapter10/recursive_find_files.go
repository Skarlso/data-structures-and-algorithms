package chapter10

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FindFiles(location string) {
	files, err := os.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			FindFiles(filepath.Join(location, file.Name()))
		}
		fmt.Println(file.Name())
	}
}
