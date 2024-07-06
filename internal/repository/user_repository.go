package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

func (repo *UserRepository) GetUsers() ([]*model.User, error) {
	rows, err := repo.DB.Query("select id, nome, usuario, createdAt from usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user = &model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{}
	query := "select id, nome, usuario, createdAt from usuario where id = $1"
	err := repo.DB.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	query := `insert into usuario (nome, usuario, senha, createdAt)
	 values ($1, $2, $3, now()) returning id`
	err := repo.DB.QueryRow(query, user.Name, user.Username, user.Password).
		Scan(&user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) UpdateUser(user *model.User) error {
	queryParts := []string{}
	values := []interface{}{}
	placeholderIndex := 1

	if user.Name != "" {
		queryParts = append(queryParts, fmt.Sprintf("nome = $%d", placeholderIndex))
		values = append(values, user.Name)
		placeholderIndex++
	}

	if user.Password != "" {
		queryParts = append(queryParts, fmt.Sprintf("senha = $%d", placeholderIndex))
		values = append(values, user.Password)
		placeholderIndex++
	}

	if len(queryParts) == 0 {
		return errors.New("nothing to update")
	}

	values = append(values, user.Id)
	query := fmt.Sprintf("UPDATE usuario SET %s WHERE id = $%d", strings.Join(queryParts, ", "), placeholderIndex)
	_, err := repo.DB.Exec(query, values...)
	return err
}

func (repo *UserRepository) DeleteUser(id int64) error {
	query := "delete from usuario where id = $1"
	_, err := repo.DB.Exec(query, id)
	return err
}
