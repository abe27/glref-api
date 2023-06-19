package routes

import (
	"github.com/abe27/vcst/api.v1/controllers"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.HelloController)

	api := c.Group("/api/v1")
	api.Get("", controllers.HelloController)

	// Auth
	c.Post("api/v1/login", controllers.LoginController)
	c.Get("api/v1/verify", controllers.VerifyTokenController)

	auth := api.Use(services.AuthorizationRequired)
	// Book
	book := auth.Group("/book")
	book.Get("", controllers.BookController)

	// whs
	whs := auth.Group("/whs")
	whs.Get("", controllers.WhsController)
	whs.Post("", controllers.WhsPostController)

	// Coor
	coor := auth.Group("/coor")
	coor.Get("", controllers.CoorController)
	// coor.Post("", controllers.CoorPostController)

	// Department
	department := auth.Group("/department")
	department.Get("", controllers.DepartmentController)
	// department.Post("", controllers.DepartmentPostController)

	// Product
	product := auth.Group("/product")
	product.Get("", controllers.ProductController)
	// product.Post("", controllers.ProductPostController)

	// Unit
	unit := auth.Group("/unit")
	unit.Get("", controllers.UnitController)
	unit.Post("", controllers.UnitPostController)

	// Glref
	glref := auth.Group("/glref")
	glref.Get("", controllers.GlrefController)
	glref.Post("", controllers.GlrefPostController)
	glref.Patch("", controllers.GlrefTransferController)

	refProd := auth.Group("/refprod")
	refProd.Get("", controllers.RefProdController)

	// Glref History
	glHistory := auth.Group("/glHistory")
	glHistory.Get("", controllers.GlrefHistoryController)
}
