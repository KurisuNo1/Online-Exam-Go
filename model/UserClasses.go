package model

import "github.com/jinzhu/gorm"

type UserClasses struct {
	gorm.Model
	UserId      int    `json:"userId"gorm:"not null"`
	ClassesId   int    `json:"classesId"gorm:"not null"`
	ClassesName string `json:"classesName"`
	UserName    string `json:"userName"gorm:"not null"`
	Position    string `json:"position"gorm:"type:varchar(110);not null"`
}
