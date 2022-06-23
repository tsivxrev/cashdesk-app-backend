package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsivxrev/cashdesk-app-backend/models"
)

func InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(&models.ErrorResponse{
		Code:   500,
		Detail: err.Error(),
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(&models.ErrorResponse{
		Code:   404,
		Detail: "Not Found",
	})
}
