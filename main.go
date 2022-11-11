package main

import (
	"fmt"
	"os"

	"github.com/ZaphCode/fiber-upload/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func handlePanics() {
	if msg := recover(); msg != nil || msg != "" {
		fmt.Println("---- Something went wrong ----")
		fmt.Println("Error:", msg)
	}
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	defer handlePanics()

	//* Init
	app := fiber.New()

	//* Settings
	LoadEnv()

	//* Routes
	app.Post("/upload-disk", controllers.UploadDiskController)
	app.Post("/upload-s3", controllers.UploadS3Controller)
	app.Post("/upload-cld", controllers.UploadCLDController)

	//* Run
	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err.Error())
	}
}

// GO AWS S3 Librarys
// go get github.com/aws/aws-sdk-go-v2
// go get github.com/aws/aws-sdk-go-v2/config
// go get github.com/aws/aws-sdk-go-v2/service/s3
// go get github.com/aws/aws-sdk-go-v2/feature/s3/manager
