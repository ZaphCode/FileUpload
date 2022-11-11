package lib

import (
	"context"
	"mime/multipart"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func getClouninaryClient() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLD_CLOUD_NAME"),
		os.Getenv("CLD_API_KEY"),
		os.Getenv("CLD_API_SECRET"),
	)
	if err != nil {
		panic(err.Error())
	}
	return cld
}

func UploadCloudinary(file multipart.File, filename string) (string, error) {
	cld := getClouninaryClient()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:   "go-fiber-upload-test",
		PublicID: filename,
	})

	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
