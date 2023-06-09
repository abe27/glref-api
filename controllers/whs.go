package controllers

import (
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func WhsController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.StoreFormula
	var Whs []models.Whs
	if c.Query("name") != "" {
		if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Where("FCCODE like ?", "%"+strings.ToUpper(c.Query("name"))+"%").Find(&Whs).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&Whs)
		}

		r.Success = true
		r.Data = &Whs
	}

	if c.Query("type") != "" {
		arr := strings.Split(c.Query("type"), ",")
		if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Where("FCCODE IN ?", arr).Find(&Whs).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&Whs)
		}

		r.Success = true
		r.Data = &Whs
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Find(&Whs).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&Whs)
	}

	r.Success = true
	r.Data = &Whs
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WhsPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Whs
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
