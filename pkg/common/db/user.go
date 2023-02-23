package db

import (
	"context"

	"github.com/pauljamescleary/gomin/pkg/common/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
}

type PostgresUserRepository struct {
	db *Database
}

func (repo PostgresUserRepository) CreateUser(user *models.User) (*models.User, error) {
	sql := `
	INSERT INTO users (id, name)
	VALUES ($1, $2)
	`
	_, err := repo.db.Conn.Exec(context.Background(), sql, user.ID, user.Name)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func NewUserRepository(db *Database) (*PostgresUserRepository, error) {
	return &PostgresUserRepository{db: db}, nil
}
