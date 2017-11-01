package handlers

import (
  "os"
  "fmt"
  utils "../utils"
  models "../models"

  "net/http"

  "github.com/labstack/echo"
  "github.com/algolia/algoliasearch-client-go/algoliasearch"
)

/* GET /search/:query */
func GetSearch(c echo.Context) error {
  dc := c.(*utils.DatabaseContext)

  var motorcycles []models.MotorcycleModel

  text := c.Param("text")

  // not safe and way to heavy
  if err := dc.Db().Where("model ILIKE ?", fmt.Sprintf("%%%s%%", text)).Limit(25).Find(&motorcycles).Error; err != nil {
    return c.JSON(500, err)
  }

	return c.JSON(http.StatusOK, motorcycles)
}

func GetSearchAlgolia(c echo.Context) error {
  client := algoliasearch.NewClient(os.Getenv("ALGOLIA_CLIENT_ID"), os.Getenv("ALGOLIA_SECRET_API_KEY"))
  index := client.InitIndex("motorcycles")

  text := c.Param("text")

  res, err := index.Search(text, nil)

  if err != nil {
    return c.JSON(500, err)
  }

  return c.JSON(http.StatusOK, res.Hits)
}
