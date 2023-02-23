package main

import (
	"github.com/pauljamescleary/gomin/pkg/common/config"
	"github.com/pauljamescleary/gomin/pkg/common/db"
	"github.com/pauljamescleary/gomin/pkg/common/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cfg, _ := config.LoadConfig()
	database := db.NewDatabase(cfg)
	userRepo, _ := db.NewUserRepository(database)
	handler := handler.NewHandler(userRepo)

	// Routes
	// e.GET("/users", api.GetAllUsers)
	e.POST("/users", handler.CreateUser)
	// e.GET("/users/:id", api.GetUser)
	// e.PUT("/users/:id", api.UpdateUser)
	// e.DELETE("/users/:id", api.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
