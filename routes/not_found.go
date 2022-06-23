package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsivxrev/cashdesk-app-backend/controllers"
)

func NotFound(app *fiber.App) {
	app.Use(controllers.NotFound)
}
