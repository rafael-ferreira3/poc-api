package dto

import (
	"time"

	"github.com/rafael-ferreira3/poc-api/internal/model"
)

type CreateUserDTO struct {
	Name     string
	Username string
	Password string
}

func (c *CreateUserDTO) ToUserModel() *model.User {
	return &model.User{
		Name:     c.Name,
		Username: c.Username,
		Password: c.Password,
	}
}

type UpdateUserDTO struct {
	Id       int64
	Name     string
	Password string
}

func (c *UpdateUserDTO) ToUserModel() *model.User {
	return &model.User{
		Id:       c.Id,
		Name:     c.Name,
		Password: c.Password,
	}
}

type UserResponseDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"nome"`
	Username  string    `json:"usuario"`
	CreatedAt time.Time `json:"created_at"`
}

func UserResponseFromUserModel(user *model.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}
