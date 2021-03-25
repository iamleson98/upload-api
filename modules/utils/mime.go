package utils

import (
	"io"
	"net/http"
)

// ContentDetector returns mimetype of a given file
func ContentDetector(reader io.Reader) (string, error) {
	buffer := make([]byte, 512)

	_, err := reader.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
