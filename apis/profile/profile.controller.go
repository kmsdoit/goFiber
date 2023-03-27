package profile

import (
	"github.com/gofiber/fiber/v2"
	"goFiber/main/models"
	"goFiber/main/services/profile"
)

func CreateProfileAPI(c *fiber.Ctx) error {
	newProfile := new(models.Profile)

	err := c.BodyParser(newProfile)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := profile.CreateProfileService(*newProfile)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}

func FileUploadAPI(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := profile.ImageFileService(file)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
