package usecases

import (
	"errors"
	"testing"
	"task_manager/Domain"
	"task_manager/Repositories"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	mockUserRepo := new(repositories.MockUserRepository)
	userUsecase := &UserUsecase{UserRepo: mockUserRepo}

	user := &domain.User{
		ID:       "1",
		Username: "testuser",
		Password: "password",
		Role:     "user",
	}

	mockUserRepo.On("CreateUser", user).Return(nil)

	err := userUsecase.RegisterUser(user)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserByUsername(t *testing.T) {
	mockUserRepo := new(repositories.MockUserRepository)
	userUsecase := &UserUsecase{UserRepo: mockUserRepo}

	user := &domain.User{
		ID:       "1",
		Username: "testuser",
		Password: "password",
		Role:     "user",
	}

	mockUserRepo.On("GetUserByUsername", "testuser").Return(user, nil)

	result, err := userUsecase.GetUserByUsername("testuser")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result)
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	mockUserRepo := new(repositories.MockUserRepository)
	userUsecase := &UserUsecase{UserRepo: mockUserRepo}

	mockUserRepo.On("GetUserByUsername", "testuser").Return(nil, errors.New("not found"))

	result, err := userUsecase.GetUserByUsername("testuser")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockUserRepo.AssertExpectations(t)
}
