package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"goFiber/main/apis/user"
	"goFiber/main/utils"
)

func SetupUserRoute(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())
	api.Get("/user/all", utils.JwtMiddleWare(), user.GetAllUserAPI)
	api.Post("/user/create", user.CreateUserAPI)
	api.Post("/user/login", user.LoginUserAPI)
	api.Get("/user/me", utils.JwtMiddleWare(), user.FindByUserEmailAPI)
	api.Post("/user/update", user.UpdateUserDataAPI)
	api.Get("/user/profiles", utils.JwtMiddleWare(), user.GetProfileByUserIdAPI)
}
