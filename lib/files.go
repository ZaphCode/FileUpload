package lib

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func ValidateFileType(file multipart.File) error {
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	mimeType := http.DetectContentType(bytes)

	if ok := stringInSlice(mimeType, []string{"image/png", "image/jpeg", "image/jpg"}); !ok {
		return errors.New("file type not supported")
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
