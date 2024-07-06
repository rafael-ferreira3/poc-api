package handler

import (
	"errors"
	"net/http"

	"github.com/rafael-ferreira3/poc-api/internal/dto"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
	"github.com/rafael-ferreira3/poc-api/internal/service"
)

var UserService = service.NewUserService()

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) error {

	usersResponse, err := UserService.GetAllUsers()
	if err != nil {
		return err
	}
	helper.WriteJson(w, http.StatusOK, usersResponse)
	return nil
}

func HandlerGetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := helper.StringToInt64(r.PathValue("id"))
	if err != nil {
		return err
	}
	userResponse, err := UserService.GetUserByID(id)
	if err != nil {
		return err
	}
	return helper.WriteJson(w, http.StatusOK, userResponse)
}

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserDTO := &dto.CreateUserDTO{}
	if err := helper.ReadRequestBody(r, createUserDTO); err != nil {
		return err
	}
	user, err := UserService.CreateUser(createUserDTO)
	if err != nil {
		return err
	}
	helper.WriteJson(w, http.StatusCreated, user)
	return nil
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) error {
	UpdateUserDTO := &dto.UpdateUserDTO{}
	if err := helper.ReadRequestBody(r, UpdateUserDTO); err != nil {
		return err
	}
	user, err := UserService.UpdateUser(UpdateUserDTO)
	if err != nil {
		return err
	}

	helper.WriteJson(w, http.StatusOK, user)

	return nil
}

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := helper.StringToInt64(r.PathValue("id"))
	if err != nil {
		return errors.New("ID inválido")
	}
	return UserService.DeleteUser(id)
}
