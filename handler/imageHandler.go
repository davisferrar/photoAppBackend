package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const publicDir = "./client/public/"

type FileDetails struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func UploadImage(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form",
		})
	}

	files := form.File["image"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No image uploaded",
		})
	}

	uploadDir := "./client/public"
	for _, file := range files {
		err = saveFile(file, uploadDir)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}
	}
	return c.SendString("Files saved")
}

func saveFile(file *multipart.FileHeader, directory string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join(directory, file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func DeleteImage(c *fiber.Ctx) error {
	fileName := c.Query("file_name")
	fileDir := "./client/public/" + fileName
	fmt.Printf(fileDir)
	err := os.Remove(fileDir)
	if err != nil {
		return err
	}
	fmt.Printf("File '%s' deleted successfully\n", fileName)
	return nil
}

func GetImage(c *fiber.Ctx) error {
	files, err := os.ReadDir(publicDir)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var fileDetailsList []FileDetails

	for _, file := range files {
		fileURL := fmt.Sprintf("http://%s%s%s", c.Hostname(), "/client/public/", file.Name())
		fileDetails := FileDetails{
			URL:  fileURL,
			Name: file.Name(),
		}
		fileDetailsList = append(fileDetailsList, fileDetails)
	}

	return c.JSON(fileDetailsList)
}
