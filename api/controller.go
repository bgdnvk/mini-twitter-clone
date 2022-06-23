package api

import (
	"goexample/services"

	"github.com/gofiber/fiber/v2"

	// "fmt"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/users", services.GetUsers)
	api.Get("/tweets", services.GetTweets)

	api.Get("/timeline/:id", services.GetTimelineTweets)
	//TODO: pagination
}
