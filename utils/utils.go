package utils

import (
	"os"
	models "../models"

	"github.com/labstack/echo"
  "github.com/jinzhu/gorm"
)

var db *gorm.DB

type DatabaseContext struct {
	echo.Context
}

func Init() {
  var err error
  db, err = gorm.Open("postgres", "host=database user=postgres DB.name=lowside sslmode=disable password=" + os.Getenv("DATABASE_PASSWORD"));

  if err != nil {
    panic("failed to connect database")
  }

	db.LogMode(true)

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"");


  // Migrate the schema
  db.AutoMigrate(&models.Motorcycle{})

  // defer db.Close()
}

func (c *DatabaseContext) Db() *gorm.DB {
  return db
}
