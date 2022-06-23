package api

import (
	"goexample/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/users", services.GetUsers)
	api.Get("/tweets", services.GetTweets)
}
