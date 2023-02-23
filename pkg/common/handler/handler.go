package handler

import (
	"github.com/pauljamescleary/gomin/pkg/common/db"
)

type Handler struct {
	UserRepo db.UserRepository
}

func NewHandler(ur db.UserRepository) *Handler {
	return &Handler{UserRepo: ur}
}
