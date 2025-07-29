package repositories

import domain "task_manager/Domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
}
