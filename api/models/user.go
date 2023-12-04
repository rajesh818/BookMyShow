package models

import (

    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name      string         `gorm:"type:varchar(255)"`
    Email     string         `gorm:"type:varchar(255)"`
    Mobile    string         `gorm:"type:varchar(255)"`
    Password  string         `gorm:"type:varchar(255)"`
}

// TableName sets the table name for the User model.
func (User) TableName() string {
    return "user"
}
