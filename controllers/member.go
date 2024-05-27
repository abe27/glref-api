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
	var isSuccess bool = true
	if err := c.BodyParser(&frm); err != nil {
		isSuccess = false
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.StoreFormula
	var emp models.Employee
	if err := db.Where("FCLOGIN=?", strings.ToUpper(frm.UserName)).Where("FCPW=?", strings.ToUpper(frm.Password)).First(&emp).Error; err != nil {
		isSuccess = false
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// Create JWT token
	auth := services.CreateToken(emp)
	r.Success = isSuccess
	r.Message = "Login Success"
	r.Data = &auth
	return c.Status(fiber.StatusOK).JSON(&r)
}

func VerifyTokenController(c *fiber.Ctx) error {
	var r models.Response
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if token == "" {
		r.Message = "Unauthorized"
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}
	_, er := services.ValidateToken(token)
	if er != nil {
		r.Message = "Token is expired"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	return c.Status(fiber.StatusOK).JSON(&r)
}
