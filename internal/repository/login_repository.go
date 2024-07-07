package repository

import (
	"database/sql"
	"errors"

	"github.com/rafael-ferreira3/poc-api/internal/database"
	"github.com/rafael-ferreira3/poc-api/internal/dto"
)

type LoginRepository struct {
	DB *sql.DB
}

func NewLoginRepository() *LoginRepository {
	return &LoginRepository{
		DB: database.DB,
	}
}

func (r *LoginRepository) Login(loginRequest *dto.LoginRequestDTO) (int64, error) {
	query := "select count(1) from usuario where usuario = $1 and senha = $2"
	var id int64 = 0
	err := r.DB.QueryRow(query, loginRequest.Username, loginRequest.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("usuário ou senha incorreto")
	}
	return id, nil
}
