package services

import (
	"echo-template/app/models"
	"echo-template/database"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: database.GetDB(),
	}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := us.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := us.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (us *UserService) CreateUser(user *models.User) error {
	if err := us.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	if err := us.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(id uint) error {
	if err := us.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
