package user

import (
	"github.com/gofiber/fiber/v2"
	"goFiber/main/models"
	"goFiber/main/services/user"
)

func CreateUserAPI(c *fiber.Ctx) error {
	newUser := new(models.User)

	err := c.BodyParser(newUser)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := user.CreateUserService(newUser.Email, newUser.Name, newUser.Password)

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

func GetAllUserAPI(c *fiber.Ctx) error {
	result, err := user.GetAllUserService()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func LoginUserAPI(c *fiber.Ctx) error {
	loginInfo := new(models.LoginRequest)

	err := c.BodyParser(loginInfo)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := user.LoginUserService(loginInfo.Email, loginInfo.Password)

	if err != nil || result == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "login request error",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}
