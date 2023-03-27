package oauth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"goFiber/main/apis/oauth"
)

func SetupOauthRoute(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())
	api.Get("/auth/google/login", oauth.GoogleLoginAPI)
	api.Get("/auth/google/callback", oauth.GoogleLoginCallBack)
}
