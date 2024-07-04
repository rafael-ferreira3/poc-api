package repository

import (
	"database/sql"

	"github.com/rafael-ferreira3/poc-api/internal/database"
	"github.com/rafael-ferreira3/poc-api/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.DB,
	}
}

func (repo *UserRepository) GetUsers() (*[]model.User, error) {
	return nil, nil
}

func (repo *UserRepository) GetUserByID() (*model.User, error) {
	return nil, nil
}

func (repo *UserRepository) CreateUser() (*model.User, error) {
	return nil, nil
}

func (repo *UserRepository) UpdateUser() (*model.User, error) {
	return nil, nil
}

func (repo *UserRepository) DeleteUser() error {
	return nil
}
