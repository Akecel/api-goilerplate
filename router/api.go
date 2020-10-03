package router

import (
	"net/http"

	controller "gechoplate/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetAPIRoutes define all apis routes.
func SetAPIRoutes(e *echo.Echo) {

	// Public group
	e.GET("/", controller.Index)

	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

	// Authentification routes
	e.POST("/login", controller.Login)
	e.POST("/regsiter", controller.Register)
	e.GET("/logout", controller.Logout)

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
}
