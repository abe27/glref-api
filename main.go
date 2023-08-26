package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres" // PGSql driver
	// Sqlite driver based on CGO

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	//// initial database
	//// github.com/mattn/go-sqlite3
	// configs.Store, err = gorm.Open(sqlite.Open("database/gorm.db"), &gorm.Config{
	// 	DisableAutomaticPing:                     true,
	// 	DisableForeignKeyConstraintWhenMigrating: false,
	// 	SkipDefaultTransaction:                   true,
	// 	NowFunc: func() time.Time {
	// 		return time.Now().Local()
	// 	},
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
	// 		SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
	// 		NoLowerCase:   false,  // skip the snake_casing of names
	// 		NameReplacer:  strings.NewReplacer("CID", "Cid"),
	// 	},
	// })
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s TimeZone=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPORT"), os.Getenv("DBPASSWORD"), os.Getenv("SSLMODE"), os.Getenv("TZNAME"))
	configs.Store, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	// // github.com/mattn/go-sqlite3
	// configs.Store, err = gorm.Open(sqlite.Open("database/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	dsnFormula := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("DB_FORMULA_MSSQL_USER"), os.Getenv("DB_FORMULA_MSSQL_PASSWORD"), os.Getenv("DB_FORMULA_MSSQL_HOST"), os.Getenv("DB_FORMULA_MSSQL_PORT"), os.Getenv("DB_FORMULA_MSSQL_DATABASE"))
	// dsnFormula := "sqlserver://sa:ADSads123@localhost:1433?database=FormulaDB"
	// fmt.Println(dsnFormula)
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
	configs.APP_LINE_TOKEN = os.Getenv("LINE_NOTFICATION")

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
