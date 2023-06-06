package main

import (
	"os"
	"time"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// initial database
	// dsnFormula := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("DB_FORMULA_MSSQL_USER"), os.Getenv("DB_FORMULA_MSSQL_PASSWORD"), os.Getenv("DB_FORMULA_MSSQL_HOST"), os.Getenv("DB_FORMULA_MSSQL_PORT"), os.Getenv("DB_FORMULA_MSSQL_DATABASE"))
	dsnFormula := "sqlserver://sa:ADSads123@localhost:1433?database=FormulaDB"
	configs.StoreFormula, err = gorm.Open(sqlserver.Open(dsnFormula), &gorm.Config{
		DisableAutomaticPing: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	configs.API_TRIGGER_URL = os.Getenv("API_TRIGGER_URL")
	configs.APP_NAME = os.Getenv("APP_NAME")
	configs.APP_DESCRIPTION = os.Getenv("APP_DESCRIPTION")

	// Auto Migration DB
	configs.SetDB()
}

func main() {
	// Create config variable
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "VCST Server API Service", // add custom server header
		AppName:       "API Version 1.0",
		BodyLimit:     10 * 1024 * 1024, // this is the default limit of 10MB
	}

	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Static("/", "./public")
	routes.SetUpRouter(app)
	app.Listen(":4040")
}
