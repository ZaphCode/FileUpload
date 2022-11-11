package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZaphCode/fiber-upload/lib"
	"github.com/gofiber/fiber/v2"
)

func UploadDiskController(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("document")

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "file not recived",
		})
	}

	file, openErr := fileHeader.Open()

	if openErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "opening the file error",
		})
	}

	defer file.Close()

	if err := lib.ValidateFileType(file); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := ctx.SaveFile(fileHeader, fmt.Sprintf("./uploads/%s", fileHeader.Filename)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "file not saved",
		})
	}

	return ctx.Status(http.StatusAccepted).JSON(fiber.Map{
		"success": true,
		"message": "file recived",
	})
}

func UploadS3Controller(ctx *fiber.Ctx) error {
	fileHeader, reciveErr := ctx.FormFile("document")

	if reciveErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "file not recived",
		})
	}

	file, openErr := fileHeader.Open()

	if openErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "error openig the file",
		})
	}

	defer file.Close()

	url, uploadErr := lib.UploadS3(file, fileHeader.Filename)

	if uploadErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": uploadErr.Error(),
		})
	}

	return ctx.Status(http.StatusAccepted).JSON(fiber.Map{
		"success": true,
		"message": "file recived",
		"key":     url,
	})
}

func UploadCLDController(ctx *fiber.Ctx) error {
	fileHeader, reciveErr := ctx.FormFile("document")

	if reciveErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "file not recived",
		})
	}

	file, openErr := fileHeader.Open()

	if openErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "error openig the file",
		})
	}

	defer file.Close()

	url, uploadErr := lib.UploadCloudinary(file, fileHeader.Filename)

	if uploadErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": uploadErr.Error(),
		})
	}

	return ctx.Status(http.StatusAccepted).JSON(fiber.Map{
		"success": true,
		"message": "file recived",
		"url":     url,
	})
}

//* -------- get protected files from s3 buckets --------

// func GetFileS3Controller(ctx *fiber.Ctx) error {
// 	key := ctx.Params("key")

// 	result, err := lib.GetFileS3(key)

// 	if err != nil {
// 		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return ctx.Status(http.StatusAccepted).JSON(fiber.Map{
// 		"success": true,
// 		"url":     result,
// 	})
// }

// func GetFilesS3Controller(ctx *fiber.Ctx) error {
// 	result, err := lib.GetFilesS3()

// 	if err != nil {
// 		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 	}

// 	return ctx.Status(http.StatusAccepted).JSON(fiber.Map{
// 		"success": true,
// 		"objects": result,
// 	})
// }
