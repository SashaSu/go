package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Пример реализации функции saveToFile
func saveToFile(fileNames []string, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, fileName := range fileNames {
		_, err = file.WriteString(fileName + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Реализовать функцию получения всех файлов из директории
func getFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil

}

// Реализовать функцию фильтрации переданных названий файлов
func filterFiles(files []string, filter string) ([]string, error) {
	var filtered []string
	if strings.HasPrefix(filter, "*."){
		extension := strings.Replace(filter, "*.", "", 1)
		for _, file := range files{
			if filepath.Ext(file) == "." + extension{
				filtered = append(filtered, file)
			}
		}
		return filtered, nil
	}

	for _, file := range files{
		if strings.Contains(file, filter) {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}

func main() {

}
