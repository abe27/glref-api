package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
)

func RefProdController(c *fiber.Ctx) error {
	var r models.Response
	var refpro []models.Refprod
	if err := configs.StoreFormula.
		Preload("Prod").
		Preload("Unit").
		Where("FCGLREF", c.Query("glref")).Find(&refpro).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Data = &refpro
	return c.Status(fiber.StatusOK).JSON(&r)
}
