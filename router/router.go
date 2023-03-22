package router

import (
	"github.com/gofiber/fiber/v2"
	"goFiber/main/apis/bookmark"
	"goFiber/main/apis/user"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! send your request")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", status)

	app.Get("/api/bookmark", bookmark.GetAllBookmarks)
	app.Post("/api/bookmark", bookmark.SaveBookmark)
	app.Post("/api/user/create", user.CreateUserAPI)
}
