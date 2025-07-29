package usecases

import (
	domain "task_manager/Domain"
	repositories "task_manager/Repositories"
)

// UserUsecase defines the business logic for users
// It depends on the UserRepository interface

type UserUsecase struct {
	UserRepo repositories.UserRepository
}

func (u *UserUsecase) RegisterUser(user *domain.User) error {
	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return u.UserRepo.GetUserByUsername(username)
}
