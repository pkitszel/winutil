package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func unzip(archive, destDir string) error {
	info, err := os.Stat(destDir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("destination %q is not a directory", destDir)
	}

	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		path := filepath.Join(destDir, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		err = func() error {
			fileReader, err := file.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()

			targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(targetFile, fileReader)
			errc := targetFile.Close()
			if err == nil {
				err = errc
			}
			return err
		}()
		if err != nil {
			return err
		}
	}

	return nil
}
