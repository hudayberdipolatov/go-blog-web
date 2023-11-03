package helpers

import (
	"log"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("views/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	path_files, err := filepath.Glob("views/" + path + "/*.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range path_files {
		files = append(files, file)
	}
	return files
}
