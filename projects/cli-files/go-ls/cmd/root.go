package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"io/ioutil"
)

func listFiles(directory string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

func readFile(filePath string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func Execute() {
	directory := "." 

	fileNames, err := listFiles(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range fileNames {
		filePath := filepath.Join(directory, fileName)

		contents, err := readFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", fileName, err)
			continue
		}

		fmt.Println(string(contents))
	}
}