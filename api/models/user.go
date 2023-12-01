package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
	Id uint `gorm:"primaryKey"`
    Name string `gorm:"column:name"`
    Email string `gorm:"column:email"`
    Mobile uint `gorm:"column:mobile"`
    Password string `gorm:"column:password"`
    CreatedAt uint `gorm:"column:created_at"`
    UpdatedAt uint `gorm:"column:updated_at"`
    DeletedAt uint `gorm:"column:deleted_at"`
}

func (User) TableName() string {
    return "user"
}