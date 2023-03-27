package profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"goFiber/main/apis/profile"
)

func SetupUserRoute(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())
	api.Post("/profile/create", profile.CreateProfileAPI)
}
