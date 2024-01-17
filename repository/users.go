package repository

import (
	"database/sql"

	"github.com/pranoyk/tiger-sightings/model"
)

type usersRepository struct {
	db *sql.DB
}

type UsersRepository interface {
	CreateUser(user model.User) error
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (ur *usersRepository) CreateUser(user model.User) error {
	_, err := ur.db.Exec("INSERT INTO USERS (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}
