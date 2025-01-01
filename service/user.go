package service

import (
	"ecommerce/repository"
	"fmt"

	"ecommerce/models"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) CreateUser(user models.User) error {
	existingUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	return u.userRepo.CreateUser(user)
}

func (u *UserService) GetUserById(userID string) (*models.User, error) {
	return u.userRepo.GetUserById(userID)
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserService) UpdateUser(id string, user models.User) error {
	exists, err := u.userRepo.UserExists(id)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	return u.userRepo.UpdateUser(id, user)
}

func (u *UserService) DeleteUser(id string) error {
	exists, err := u.userRepo.UserExists(id)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	return u.userRepo.DeleteUser(id)
}
