package configs

import (
	"github.com/abe27/vcst/api.v1/models"
	"gorm.io/gorm"
)

var (
	Store           *gorm.DB
	StoreVCST       *gorm.DB
	StoreFormula    *gorm.DB
	StoreITC        *gorm.DB
	StoreBSV        *gorm.DB
	StoreAAA        *gorm.DB
	API_TRIGGER_URL string
	APP_NAME        string
	APP_DESCRIPTION string
	APP_LINE_TOKEN  string
)

func SetDB() {
	if !Store.Migrator().HasTable(&models.GlrefHistory{}) {
		Store.AutoMigrate(&models.GlrefHistory{})
	}

	// if !StoreFormula.Migrator().HasTable(&models.Product{}) {
	// 	StoreFormula.AutoMigrate(&models.Product{})
	// }

	// if !StoreFormula.Migrator().HasTable(&models.ProductType{}) {
	// 	StoreFormula.AutoMigrate(&models.ProductType{})
	// }
}
