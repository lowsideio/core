package handlers

import (
	utils "github.com/lowsideio/core/utils"
	models "github.com/lowsideio/core/models"

	"fmt"

	"net/http"

	"github.com/labstack/echo"
)

/* GET /search/:query */
func GetSearch(c echo.Context) error {
	dc := c.(*utils.RequestContext)

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
	ac := c.(*utils.RequestContext)

	text := c.Param("text")

	res, err := ac.AlgoliaIndex().Search(text, nil)

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(http.StatusOK, res.Hits)
}
