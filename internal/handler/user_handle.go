package handler

import (
	"net/http"
	"strconv"

	"github.com/rafael-ferreira3/poc-api/internal/model"
	"github.com/rafael-ferreira3/poc-api/internal/util"
)

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HandlerGetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return err
	}
	user := model.User{Id: id, Name: "Nome"}
	return util.WriteJson(w, http.StatusOK, user)
}

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
