package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/rafael-ferreira3/poc-api/internal/dto"
	"github.com/rafael-ferreira3/poc-api/internal/service"
	"github.com/rafael-ferreira3/poc-api/internal/util"
)

var UserService = service.NewUserService()

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) error {

	usersResponse, err := UserService.GetAllUsers()
	if err != nil {
		return err
	}
	util.WriteJson(w, http.StatusOK, usersResponse)
	return nil
}

func HandlerGetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 0, 64)
	if err != nil {
		return err
	}
	userResponse, err := UserService.GetUserByID(id)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, userResponse)
}

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserDTO := &dto.CreateUserDTO{}
	if err := json.NewDecoder(r.Body).Decode(createUserDTO); err != nil {
		return err
	}
	user, err := UserService.CreateUser(createUserDTO)
	if err != nil {
		return err
	}
	util.WriteJson(w, http.StatusCreated, user)
	return nil
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) error {
	UpdateUserDTO := &dto.UpdateUserDTO{}
	if err := json.NewDecoder(r.Body).Decode(UpdateUserDTO); err != nil {
		return err
	}
	user, err := UserService.UpdateUser(UpdateUserDTO)
	if err != nil {
		return err
	}

	util.WriteJson(w, http.StatusOK, user)

	return nil
}

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 0, 64)
	if err != nil {
		return errors.New("ID inválido")
	}
	return UserService.DeleteUser(id)
}
