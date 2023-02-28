package router

import (
	"github.com/pauljamescleary/gomin/pkg/common/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(handler *handler.Handler) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// e.GET("/users", api.GetAllUsers)
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	// e.PUT("/users/:id", api.UpdateUser)
	// e.DELETE("/users/:id", api.DeleteUser)
	return e
}
