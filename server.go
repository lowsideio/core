package main

import (
  handlers "./handlers"
  utils "./utils"

  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func init() {
  utils.Init()
}

func main() {

  e := echo.New()

  e.Debug = true

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:3000"},
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
  }))

  e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      cc := &utils.DatabaseContext{c}
      return h(cc)
    }
  })

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  e.GET("/motorcycles/:id", handlers.GetMotorcycle)
  e.PUT("/motorcycles/:id", handlers.PutMotorcycle)
  e.DELETE("/motorcycles/:id", handlers.DeleteMotorcycle)
  e.OPTIONS("/motorcycles/:id", handlers.MotorcyclesOptions)

  e.GET("/motorcycles", handlers.GetMotorcycles)
  e.POST("/motorcycles", handlers.PostMotorcycles)
  e.OPTIONS("/motorcycles", handlers.MotorcyclesOptions)

  e.Logger.Fatal(e.Start(":1323"))
}
