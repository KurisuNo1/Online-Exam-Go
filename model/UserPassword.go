package model

import "github.com/jinzhu/gorm"

type UserPassword struct {
	gorm.Model
	UserId   string `json:"userId"gorm:"type:varchar(110);not null"`
	Password string `json:"password"gorm:"type:varchar(110);not null"`
}
