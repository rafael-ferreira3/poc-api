package service

import (
	"github.com/rafael-ferreira3/poc-api/internal/dto"
	"github.com/rafael-ferreira3/poc-api/internal/repository"
)

type LoginService struct {
	LoginRepository *repository.LoginRepository
	UserService     *UserService
	TokenService    *TokenService
}

func NewLoginService() *LoginService {
	return &LoginService{
		LoginRepository: repository.NewLoginRepository(),
		UserService:     NewUserService(),
		TokenService:    NewTokenService(),
	}
}

func (l *LoginService) Login(loginRequestDTO *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error) {
	id, err := l.LoginRepository.Login(loginRequestDTO)
	if err != nil {
		return nil, err
	}
	userDTO, err := l.UserService.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	token, err := l.TokenService.CreateJWTToken(userDTO.Username)
	if err != nil {
		return nil, err
	}
	loginResponse := &dto.LoginResponseDTO{
		ID:          userDTO.ID,
		Name:        userDTO.Name,
		AccessToken: token,
	}
	return loginResponse, nil
}
