package usecases

import (
	domain "task_manager/Domain"
	repositories "task_manager/Repositories"
)

// TaskUsecase defines the business logic for tasks
// It depends on the TaskRepository interface

type TaskUsecase struct {
	TaskRepo repositories.TaskRepository
}

func (u *TaskUsecase) CreateTask(task *domain.Task) error {
	return u.TaskRepo.CreateTask(task)
}

func (u *TaskUsecase) GetTaskByID(id string) (*domain.Task, error) {
	return u.TaskRepo.GetTaskByID(id)
}

func (u *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
	return u.TaskRepo.GetAllTasks()
}

func (u *TaskUsecase) UpdateTask(task *domain.Task) error {
	return u.TaskRepo.UpdateTask(task)
}

func (u *TaskUsecase) DeleteTask(id string) error {
	return u.TaskRepo.DeleteTask(id)
}
