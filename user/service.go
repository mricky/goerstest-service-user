package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserInput(input RegisterUserInput) (Users, error)
	Login(input LoginInput) (Users, error)
	GetUserById(input UserInput) (Users, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserById(input UserInput) (Users, error) {
	user, err := s.repository.FindById(input.ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("NO USER FIND IN EMAIL")
	}

	if err != nil {
		return user, nil
	}
	return user, nil
}
func (s *service) RegisterUserInput(input RegisterUserInput) (Users, error) {
	user := Users{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "User"
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *service) Login(input LoginInput) (Users, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email, password)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("NO USER FIND IN EMAIL")
	}

	if err != nil {
		return user, nil
	}
	return user, nil
}
