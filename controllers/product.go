package controllers

import (
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func ProductController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.StoreFormula
	var Product []models.Product
	if c.Query("name") != "" {
		if err := db.Scopes(services.Paginate(c)).Where("FCCODE like ?", "%"+strings.ToUpper(c.Query("name"))+"%").Find(&Product).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&Product)
		}

		r.Success = true
		r.Data = &Product
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	arr := strings.Split(c.Query("type"), ",")
	if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Where("FCTYPE IN ?", arr).Find(&Product).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&Product)
	}

	r.Success = true
	r.Data = &Product
	return c.Status(fiber.StatusOK).JSON(&r)
}
