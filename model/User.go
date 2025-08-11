package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"size:255;not null"`
	Phone     string `json:"phone" gorm:"type:varchar(20);not null;unique"`
	Password  string `json:"password" gorm:"size:255;not null"`
	Email     string `json:"email" gorm:"size:255;not null"`
	Sex       string `json:"sex" gorm:"type:varchar(110);not null"`
	PhotoName string `json:"photoName" gorm:"type:varchar(110);not null"`
	Photo     string `json:"photo"`
	Role      string `json:"role" gorm:"type:varchar(110);not null"`
	Work      string `json:"work" gorm:"type:varchar(110)"`
}
