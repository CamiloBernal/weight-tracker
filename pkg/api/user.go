package api

import (
	"errors"
	"log"
	"strings"
)

// UserService contains the methods of the user service
type UserService interface {
	New(user NewUserRequest) error
}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	CreateUser(NewUserRequest) error
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{storage: userRepo}
}

func (u *userService) New(user NewUserRequest) error {

	log.Printf("The user data: %v", user)

	if user.Email == "" {
		return errors.New("user service - email is required")
	}
	if user.Name == "" {
		return errors.New("user service - name is required")
	}

	if user.WeightGoal == "" {
		return errors.New("user service - weight goal is required")
	}

	user.Name = strings.ToLower(user.Name)
	user.Email = strings.ToLower(user.Email)

	err := u.storage.CreateUser(user)

	if err != nil {
		return err
	}
	return nil
}
