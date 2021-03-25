package zipper

import (
	"archive/zip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func ZipExtractHandler(zipFile multipart.File, fileHeader *multipart.FileHeader, ch chan<- error) {
	defer func() {
		close(ch)
	}()

	reader, err := zip.NewReader(zipFile, fileHeader.Size)
	if err != nil {
		ch <- err
		return
	}

	for _, file := range reader.File {
		fPath := filepath.Join("dest", file.Name)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fPath, os.ModePerm)
			if err != nil {
				ch <- err
				return
			}
			continue
		}

		// Make file
		if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			ch <- err
			return
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			ch <- err
			return
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			ch <- err
			return
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			ch <- err
			return
		}
	}
}
