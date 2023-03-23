package bookmark

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"goFiber/main/apis/bookmark"
)

func SetupBookmarkRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	api.Get("/bookmark/all", bookmark.GetAllBookmarks)
	api.Post("/bookmark/create", bookmark.SaveBookmark)
}
