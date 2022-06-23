package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsivxrev/cashdesk/controllers"
)

func NotFound(app *fiber.App) {
	app.Use(controllers.NotFound)
}
