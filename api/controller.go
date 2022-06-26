package api

import (
	"goexample/services"

	"github.com/gofiber/fiber/v2"
	// "fmt"
)

//http://localhost:3000/api/timeline/1

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/users", services.GetUsers)
	api.Get("/tweets", services.GetTweets)

	api.Get("/timeline/:id", services.GetTimelineTweets)
	//TODO: pagination

	//users with age
	api.Get("/users/age", services.GetUsersByAgeAsc)
}
