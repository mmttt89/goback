package service

import (
	"goback/src/model"
	"goback/src/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {

	user, err := s.userRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.userRepository.GetAllUsers()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) AddUser(newUser model.User) (*model.User, error) {
	user, err := s.userRepository.AddUser(newUser)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *UserService) UpdateUser(UpdateUser model.User) (bool, error) {
	isUpdated := false
	user, err := s.userRepository.FindByEmail(UpdateUser.Email)

	if err != nil {
		return false, err
	}
	if user != nil {
		isUpdated, updateErr := s.userRepository.UpdateUser(UpdateUser)
		if updateErr != nil || !isUpdated {
			return false, updateErr
		}
		return isUpdated, nil
	}

	return isUpdated, nil
}
