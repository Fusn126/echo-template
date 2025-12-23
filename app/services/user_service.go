package services

import (
	"echo-template/app/models"
	"echo-template/database"
	"echo-template/utils"
	"errors"

	"gorm.io/gorm"
)

// 确保 UserService 实现了 UserServiceInterface
var _ UserServiceInterface = (*UserService)(nil)

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
		return nil, utils.ErrInternal("查询用户列表失败", err)
	}
	return users, nil
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := us.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound("用户不存在")
		}
		return nil, utils.ErrInternal("查询用户失败", err)
	}
	return &user, nil
}

func (us *UserService) CreateUser(user *models.User) error {
	if err := us.db.Create(user).Error; err != nil {
		return utils.ErrInternal("创建用户失败", err)
	}
	return nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	if err := us.db.Save(user).Error; err != nil {
		return utils.ErrInternal("更新用户失败", err)
	}
	return nil
}

func (us *UserService) DeleteUser(id uint) error {
	if err := us.db.Delete(&models.User{}, id).Error; err != nil {
		return utils.ErrInternal("删除用户失败", err)
	}
	return nil
}
