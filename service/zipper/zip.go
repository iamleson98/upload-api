package zipper

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/leminhson2398/zipper/pkg/logger"
)

// ZipExtractHandler performs extract uploaded zip files
func ZipExtractHandler(zipFile multipart.File, ch chan<- error) {

	tempFile, err := ioutil.TempFile("media", "upload-*.zip")
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		ch <- err
	}
	defer tempFile.Close()
	defer zipFile.Close()

	_, err = io.Copy(tempFile, zipFile)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		ch <- err
	}

	// extract
	err = Unzip(tempFile.Name())
	if err != nil {
		ch <- err
	}

	ch <- nil
	close(ch)
}

func Unzip(src string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		fPath := filepath.Join("dest", file.Name)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fPath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// Make file
		if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
