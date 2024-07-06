package service

import (
	"errors"

	"github.com/rafael-ferreira3/poc-api/internal/dto"
	"github.com/rafael-ferreira3/poc-api/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

func (u *UserService) GetAllUsers() ([]*dto.UserResponseDTO, error) {
	users, err := u.UserRepository.GetUsers()
	var usersResponse []*dto.UserResponseDTO
	for _, user := range users {
		usersResponse = append(usersResponse, dto.UserResponseFromUserModel(user))
	}
	return usersResponse, err
}

func (u *UserService) GetUserByID(id int64) (*dto.UserResponseDTO, error) {
	user, err := u.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	userResponse := dto.UserResponseFromUserModel(user)
	return userResponse, nil
}

func (u *UserService) CreateUser(createUserDTO *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	userToCreate := createUserDTO.ToUserModel()
	user, err := u.UserRepository.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	userResonde := dto.UserResponseFromUserModel(user)
	return userResonde, nil
}

func (u *UserService) UpdateUser(updateUserDTO *dto.UpdateUserDTO) (*dto.UserResponseDTO, error) {
	if updateUserDTO.Id == 1 {
		return nil, errors.New("não é possível alterar o usuário admin")
	}
	user := updateUserDTO.ToUserModel()
	err := u.UserRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	user, err = u.UserRepository.GetUserByID(updateUserDTO.Id)
	if err != nil {
		return nil, err
	}

	return dto.UserResponseFromUserModel(user), nil
}

func (u *UserService) DeleteUser(id int64) error {
	if id == 1 {
		return errors.New("não é possível deletar o usuário admin")
	}
	_, err := u.UserRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	return u.UserRepository.DeleteUser(id)
}
