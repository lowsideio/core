package handlers

import (
  "fmt"

  utils "../utils"
	models "../models"

	"net/http"

	"github.com/labstack/echo"
)

func GetMotorcycle(c echo.Context) error {
  dc := c.(*utils.RequestContext)

  var motorcycle models.Motorcycle

	id := c.Param("id")

  if err := dc.Db().Where("id = ?", id).First(&motorcycle).Error; err != nil {
    return c.JSON(404, err)
  }

	return c.JSON(http.StatusOK, motorcycle)
}

func GetMotorcycles(c echo.Context) error {
  dc := c.(*utils.RequestContext)

  var motorcycles []models.Motorcycle
  var count int64

  if err := dc.Db().Find(&motorcycles).Error; err != nil {
    return c.JSON(404, err)
  }

  if err := dc.Db().Table("motorcycles").Count(&count).Error; err != nil {
    return c.JSON(500, err)
  }

  c.Response().Header().Set("X-Total-Count", fmt.Sprintf("%d", count));
  c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Range");
  c.Response().Header().Set("Content-Type", "application/json");
  c.Response().Header().Set("Content-Range", fmt.Sprintf("items 0-%d/%d", count, count));

  return c.JSON(200, motorcycles)
}

func PutMotorcycle(c echo.Context) error {
  dc := c.(*utils.RequestContext)

  m := &models.Motorcycle{}
  newVersion := &models.Motorcycle{}

  if err := c.Bind(m); err != nil {
    return err
  }

  newVersion.ID = m.ID;

  dc.Db().Model(&newVersion).Updates(&m);

  return c.JSON(200, newVersion)
}


func PostMotorcycles(c echo.Context) error {
  dc := c.(*utils.RequestContext)
  m := &models.Motorcycle{}

  if err := c.Bind(m); err != nil {
    return err
  }

  dc.Db().Create(&m)

  return c.JSON(200, m)
}


func DeleteMotorcycle(c echo.Context) error {
  dc := c.(*utils.RequestContext)

  var motorcycle models.Motorcycle

	id := c.Param("id")

  if err := dc.Db().Where("id = ?", id).First(&motorcycle).Error; err != nil {
    return c.JSON(404, err)
  }

  dc.Db().Delete(&motorcycle)

	return c.JSON(http.StatusOK, motorcycle)
}

func MotorcyclesOptions(c echo.Context) error {
  c.Response().WriteHeader(http.StatusNoContent)
  return nil
}
