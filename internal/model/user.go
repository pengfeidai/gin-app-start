package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primarykey" json:"id" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdateAt  time.Time      `json:"update_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
	Username  string         `gorm:"size:64;uniqueIndex;not null" json:"username" binding:"required" example:"john_doe"`
	Email     string         `gorm:"size:128;uniqueIndex" json:"email" example:"john@example.com"`
	Phone     string         `gorm:"size:32;uniqueIndex" json:"phone" example:"13800138000"`
	Password  string         `gorm:"size:128;not null" json:"-" swaggerignore:"true"`
	Salt      string         `gorm:"size:32;not null" json:"-" swaggerignore:"true"`
	Avatar    string         `gorm:"size:256" json:"avatar" example:"https://example.com/avatar.jpg"`
	Status    int8           `gorm:"default:1;not null" json:"status" example:"1"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdateAt = time.Now()
	if u.Status == 0 {
		u.Status = 1
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdateAt = time.Now()
	return nil
}

