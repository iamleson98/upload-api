package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leminhson2398/zipper/modules/setting"
	"github.com/leminhson2398/zipper/modules/utils"
	"github.com/leminhson2398/zipper/pkg/logger"
	"github.com/leminhson2398/zipper/pkg/token"
	"github.com/leminhson2398/zipper/service/zipper"
)

// TokenGenerator returns access token
func TokenGenerator(w http.ResponseWriter, r *http.Request) {
	tokn, err := token.NewAccessToken(setting.DefaultTokenExpireDuration)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokn))
}

// ZipUploadHandler handles zipped binary files uploaded by user
func ZipUploadHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20) // at most 10MB allowed

	// get fuploaded .zip file with key
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Logger.Error().Msg(err.Error())
		return
	}
	defer file.Close()

	if mime, err := utils.ContentDetector(file); err != nil || mime != "application/zip" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("only .zip files are accepted\n"))
		return
	}

	logger.Logger.Info().Msg(fmt.Sprintf("Got file name: %s size: %d", fileHeader.Filename, fileHeader.Size))

	ch := make(chan error)
	// the zip extract function should take less than 3 mins
	// otherwise, function will return
	timer := time.After(3 * time.Minute)

	go zipper.ZipExtractHandler(file, fileHeader, ch)

	for {
		select {
		case err := <-ch:
			if err != nil {
				logger.Logger.Error().Msg(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				logger.Logger.Debug().Msg("upzipped successfully")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("upzipped successfully\n"))
			}
			return
		case <-timer:
			logger.Logger.Error().Msg("time out")
			w.WriteHeader(http.StatusRequestTimeout)
			return
		}
	}
}
