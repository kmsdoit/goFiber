package router

import (
	"github.com/gofiber/fiber/v2"
	"goFiber/main/router/bookmark"
	"goFiber/main/router/oauth"
	"goFiber/main/router/profile"
	"goFiber/main/router/user"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! send your request")
}
func SetupRoutes(app *fiber.App) {
	app.Get("/", status)
	user.SetupUserRoute(app)
	bookmark.SetupBookmarkRoutes(app)
	oauth.SetupOauthRoute(app)
	profile.SetupUserRoute(app)
}
