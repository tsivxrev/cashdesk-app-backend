package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Static(app *fiber.App) {
	app.Static("/static", "./static")
}
