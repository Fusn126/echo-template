package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
// @Description 用户信息
type User struct {
	ID        uint           `json:"id" example:"1" gorm:"primarykey"`                                    // 用户ID
	CreatedAt time.Time      `json:"created_at" example:"2024-01-01T00:00:00Z"`                           // 创建时间
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-01T00:00:00Z"`                           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                                                       // 删除时间（不返回）
	
	Username string `json:"username" example:"john_doe" gorm:"uniqueIndex;not null" binding:"required"`   // 用户名
	Email    string `json:"email" example:"john@example.com" gorm:"uniqueIndex;not null" binding:"required,email"` // 邮箱
	Password string `json:"-" gorm:"not null"`                                                           // 密码（不返回）
	Name     string `json:"name" example:"John Doe"`                                                       // 姓名
}

func (User) TableName() string {
	return "users"
}

