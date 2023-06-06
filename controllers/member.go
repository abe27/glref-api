package controllers

import (
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmLogin
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.StoreFormula
	var emp models.Employee
	if err := db.Where("FCLOGIN=?", strings.ToUpper(frm.UserName)).Where("FCPW=?", strings.ToUpper(frm.Password)).First(&emp).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// Create JWT token
	auth := services.CreateToken(emp)
	r.Message = "Login Success"
	r.Data = &auth
	return c.Status(fiber.StatusOK).JSON(&r)
}
