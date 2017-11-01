package handlers

import (
  "fmt"
  utils "../utils"
	models "../models"

	"net/http"

	"github.com/labstack/echo"
)

/* GET /search/:query */
func GetSearch(c echo.Context) error {
  dc := c.(*utils.DatabaseContext)

  var motorcycles []models.Motorcycle

	text := c.Param("text")

  // not safe and way to heavy
  if err := dc.Db().Select("brand, model, power, cc, category").Where("model ILIKE ?", fmt.Sprintf("%%%s%%", text)).Limit(25).Find(&motorcycles).Error; err != nil {
    return c.JSON(500, err)
  }

	return c.JSON(http.StatusOK, motorcycles)
}
