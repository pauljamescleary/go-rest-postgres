package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name"`
}
