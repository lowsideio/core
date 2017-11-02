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

  var motorcycles []models.MotorcycleModel

  text := c.Param("text")

  // not safe and way too heavy
  if err := dc.Db().Where("model ILIKE ?", fmt.Sprintf("%%%s%%", text)).Limit(25).Find(&motorcycles).Error; err != nil {
    return c.JSON(500, err)
  }

	return c.JSON(http.StatusOK, motorcycles)
}

/* GET /search-algolia/:text */
func GetSearchAlgolia(c echo.Context) error {
  ac := c.(*utils.AlgoliaContext)

  text := c.Param("text")

  res, err := ac.AlgoliaIndex().Search(text, nil)

  if err != nil {
    return c.JSON(500, err)
  }

  return c.JSON(http.StatusOK, res.Hits)
}
