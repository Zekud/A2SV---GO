package repositories

import domain "task_manager/Domain"

type TaskRepository interface {
	CreateTask(task *domain.Task) error
	GetTaskByID(id string) (*domain.Task, error)
	GetAllTasks() ([]*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
}
