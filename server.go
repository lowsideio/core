package main

import (
	"os"

	handlers "github.com/lowsideio/core/handlers"
	utils "github.com/lowsideio/core/utils"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	utils.Init()
	utils.InitSearch()
}

func main() {

	e := echo.New()

	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CORS_URL")},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &utils.RequestContext{c}
			return h(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/motorcycles/slug/:slug", handlers.GetMotorcycleByslug)
	e.GET("/motorcycles-specs/:id", handlers.GetMotorcycleSpecsByModelId)

	e.GET("/motorcycles/:id", handlers.GetMotorcycle)
	e.PUT("/motorcycles/:id", handlers.PutMotorcycle)
	e.DELETE("/motorcycles/:id", handlers.DeleteMotorcycle)
	e.OPTIONS("/motorcycles/:id", handlers.MotorcyclesOptions)

	// e.GET("/motorcycles", handlers.GetMotorcycles)
	e.POST("/motorcycles", handlers.PostMotorcycles)
	e.OPTIONS("/motorcycles", handlers.MotorcyclesOptions)

	e.GET("/search/:text", handlers.GetSearch)
	e.GET("/search-algolia/:text", handlers.GetSearchAlgolia)

	e.Logger.Fatal(e.Start(":1323"))
}
