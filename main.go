package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tsivxrev/cashdesk/controllers"
	"github.com/tsivxrev/cashdesk/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: controllers.InternalServerError,
	})

	app.Use(recover.New())

	routes.Static(app)   // /static
	routes.Entries(app)  // /entries
	routes.NotFound(app) // Not Found handler

	app.Listen(":3001")
}
