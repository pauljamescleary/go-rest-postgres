package handler

import (
	"fmt"

	"github.com/pauljamescleary/gomin/pkg/common/config"
	"github.com/pauljamescleary/gomin/pkg/common/db"
)

type Handler struct {
	UserRepo db.UserRepository
}

func NewHandler(ur db.UserRepository) *Handler {
	return &Handler{UserRepo: ur}
}

func LoadHandler(configPath *string) *Handler {
	cfg, _ := config.LoadConfig(configPath)
	return LoadHandlerFromConfig(cfg)
}

func LoadHandlerFromConfig(cfg config.Config) *Handler {
	fmt.Printf("*** DB URL %s", cfg.DbUrl)

	database := db.NewDatabase(cfg)
	userRepo, _ := db.NewUserRepository(database)
	handler := NewHandler(userRepo)

	return handler
}
