package model

import "github.com/jinzhu/gorm"

type Subjects struct {
	gorm.Model
	SubjectsId   int    `gorm:"not null"`
	SubjectsName string `gorm:"type:varchar(110);not null"`
	CreateId     int    `gorm:"not null"`
	CreateName   string `gorm:"type:varchar(110);not null"`
}
