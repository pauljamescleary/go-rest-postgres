package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pauljamescleary/gomin/pkg/common/models"
)

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	u.ID = uuid.New()
	newUser, err := h.UserRepo.CreateUser(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.UserRepo.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, id)
	}

	return c.JSON(http.StatusOK, user)
}
