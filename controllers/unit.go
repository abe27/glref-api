package controllers

import (
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func UnitController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.StoreFormula
	var Unit []models.Unit
	if c.Query("name") != "" {
		if err := db.Scopes(services.Paginate(c)).Where("FCCODE like ?", "%"+strings.ToUpper(c.Query("name"))+"%").Find(&Unit).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&Unit)
		}

		r.Success = true
		r.Data = &Unit
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Find(&Unit).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&Unit)
	}

	r.Success = true
	r.Data = &Unit
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UnitPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Unit
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	if err := configs.StoreFormula.Create(&frm).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Success = true
	r.Data = &frm
	return c.Status(fiber.StatusCreated).JSON(&r)
}
