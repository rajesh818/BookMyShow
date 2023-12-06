package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	City    string `gorm:"type:varchar(255)"`
	State   string `gorm:"type:varchar(255)"`
	Pincode int    `gorm:"type:int"`
}

// TableName sets the table name for the User model.
func (Location) TableName() string {
	return "location"
}
