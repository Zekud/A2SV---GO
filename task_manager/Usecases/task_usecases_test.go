package usecases

import (
	"errors"
	"testing"
	"time"
	"task_manager/Domain"
	"task_manager/Repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockTaskRepo := new(repositories.MockTaskRepository)
	taskUsecase := &TaskUsecase{TaskRepo: mockTaskRepo}

	task := &domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}

	mockTaskRepo.On("CreateTask", task).Return(nil)

	err := taskUsecase.CreateTask(task)

	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	mockTaskRepo := new(repositories.MockTaskRepository)
	taskUsecase := &TaskUsecase{TaskRepo: mockTaskRepo}

	task := &domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}

	mockTaskRepo.On("GetTaskByID", "1").Return(task, nil)

	result, err := taskUsecase.GetTaskByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, task, result)
	mockTaskRepo.AssertExpectations(t)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	mockTaskRepo := new(repositories.MockTaskRepository)
	taskUsecase := &TaskUsecase{TaskRepo: mockTaskRepo}

	mockTaskRepo.On("GetTaskByID", "1").Return(nil, errors.New("not found"))

	result, err := taskUsecase.GetTaskByID("1")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockTaskRepo.AssertExpectations(t)
}

func TestGetAllTasks(t *testing.T) {
	mockTaskRepo := new(repositories.MockTaskRepository)
	taskUsecase := &TaskUsecase{TaskRepo: mockTaskRepo}

	tasks := []*domain.Task{
		{
			ID:          "1",
			Title:       "Test Task 1",
			Description: "Test Description 1",
			DueDate:     time.Now(),
			Status:      "Pending",
		},
		{
			ID:          "2",
			Title:       "Test Task 2",
			Description: "Test Description 2",
			DueDate:     time.Now(),
			Status:      "Done",
		},
	}

	mockTaskRepo.On("GetAllTasks").Return(tasks, nil)

	result, err := taskUsecase.GetAllTasks()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, tasks, result)
	mockTaskRepo.AssertExpectations(t)
}
