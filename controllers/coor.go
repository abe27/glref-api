package controllers

import (
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func CoorController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.StoreFormula
	var coor []models.Coor
	if c.Query("name") != "" {
		if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Where("FCCODE like ?", "%"+strings.ToUpper(c.Query("name"))+"%").Find(&coor).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&coor)
		}

		r.Success = true
		r.Data = &coor
	}

	if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Find(&coor).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&coor)
	}

	r.Success = true
	r.Data = &coor
	return c.Status(fiber.StatusOK).JSON(&r)
}
