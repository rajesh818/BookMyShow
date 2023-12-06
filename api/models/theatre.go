package models

import (
	"gorm.io/gorm"
)

type Theatre struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255)"`
	OwnerEmail   string `gorm:"type:varchar(255);column:owner_email"`
	OwnerMobile   string `gorm:"type:varchar(255);column:owner_mobile"`
	LocationID int `gorm:"type:int;column:location_id"`
	Location Location `gorm:"foreignKey:LocationID"`
}

// TableName sets the table name for the User model.
func (Theatre) TableName() string {
	return "theatre"
}
