package main

import (
	"net/http"

	"gitlab.com/AnhellO/go-experiments/samples/apis/restful-stock-api-echo/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	// Logger
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	// Price endpoint
	e.POST("/price", controllers.GrabPrice)
	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}
