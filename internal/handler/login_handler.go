package handler

import (
	"net/http"

	"github.com/rafael-ferreira3/poc-api/internal/dto"
	"github.com/rafael-ferreira3/poc-api/internal/helper"
	"github.com/rafael-ferreira3/poc-api/internal/service"
)

var loginService = service.NewLoginService()

func HandlerLogin(w http.ResponseWriter, r *http.Request) error {
	loginRequestDTO := &dto.LoginRequestDTO{}
	if err := helper.ReadRequestBody(r, loginRequestDTO); err != nil {
		return nil
	}
	loginResponseDTO, err := loginService.Login(loginRequestDTO)
	if err != nil {
		return err
	}
	helper.WriteJson(w, http.StatusOK, loginResponseDTO)
	return nil
}
