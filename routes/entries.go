package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsivxrev/cashdesk/controllers"
)

func Entries(app *fiber.App) {
	entries := app.Group("/entries")

	entries.Post("/", controllers.CreateEntry)
	entries.Get("/", controllers.GetAllEntries)
	entries.Get("/:entryId", controllers.GetEntry)
	entries.Put("/:entryId", controllers.EditEntry)
}
