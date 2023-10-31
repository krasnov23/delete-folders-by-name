package service

import (
	"log"
	"os"
	"path/filepath"
)

func CleanAllFilesFromFolder(folderName string) {

	folders := []string{"names", "data"}

	dir := "./" + folderName

	for _, folder := range folders {
		if folder == folderName {
			d, err := os.Open(dir)
			if err != nil {
				log.Fatalf("Unable to open directory: %v", err)
			}
			defer d.Close()

			// Читает каждый из файлов находящихся внутри указанной папки
			names, err := d.Readdirnames(-1)
			if err != nil {
				log.Fatalf("Unable to get dir names: %v", err)
			}

			for _, name := range names {
				os.RemoveAll(filepath.Join(dir, name)) // remove files in directory
			}

		}
	}

}
