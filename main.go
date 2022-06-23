package main

import (
	"github.com/gofiber/fiber/v2"
	"goexample/database"
	"log"

	"goexample/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api.SetupRoutes(app)


    log.Fatal(app.Listen(":3000"))

}
