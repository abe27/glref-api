package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
)

func HelloController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Hello World!"
	return c.Status(fiber.StatusOK).JSON(&r)
}
