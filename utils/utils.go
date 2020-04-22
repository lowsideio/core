package utils

import (
	models "github.com/lowsideio/core/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

var db *gorm.DB

type RequestContext struct {
	echo.Context
}

func Init() {
	var err error
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("DATABASE_HOST")+
			" user="+os.Getenv("DATABASE_USER")+
			" dbname="+os.Getenv("DATABASE_NAME")+
			" sslmode=disable"+
			" password="+os.Getenv("DATABASE_PASSWORD"))

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Migrate the schema
	db.AutoMigrate(&models.Motorcycle{})
	db.AutoMigrate(&models.MotorcycleModel{})

	// defer db.Close()
}

func (c *RequestContext) Db() *gorm.DB {
	return db
}
