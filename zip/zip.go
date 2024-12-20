package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Folder(folderPath, zipPath, fileName string) {
	currentDate := time.Now().Format("2006-01-02")
	zp := filepath.Join(zipPath, fmt.Sprintf("%v_%v.zip", currentDate, fileName))
	zipFile, err := os.Create(zp)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the directory
	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Create relative path for zip
		relPath, err := filepath.Rel(folderPath, path)
		if err != nil {
			return err
		}

		// Create zip file entry
		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// Open and copy the file contents
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipFile, file)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}
