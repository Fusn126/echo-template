package services

import "echo-template/app/models"

// UserServiceInterface 用户服务接口
type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

