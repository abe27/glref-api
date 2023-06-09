package controllers

import (
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func BookController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.StoreFormula
	var book []models.Booking

	bookType := strings.Split(c.Query("type"), ",")
	if c.Query("name") != "" {
		if err := db.Scopes(services.Paginate(c)).Where("FCREFTYPE", bookType).Where("FCCODE like ?", "%"+strings.ToUpper(c.Query("name"))+"%").Find(&book).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&book)
		}

		r.Success = true
		r.Data = &book
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := db.Order("FCCODE").Where("FCREFTYPE IN ?", bookType).Find(&book).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&book)
	}

	r.Success = true
	r.Data = &book
	return c.Status(fiber.StatusOK).JSON(&r)
}
