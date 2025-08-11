package model

import "github.com/jinzhu/gorm"

type Classes struct {
	gorm.Model
	ClassesId    int    `json:"classesId"gorm:"not null"`
	ClassesName  string `json:"classesName"gorm:"type:varchar(110);not null"`
	Introduction string `json:"introduction"gorm:"type:varchar(110);not null"`
	PeopleNum    int    `json:"peopleNum"gorm:"not null"`
	CreatorName  string `json:"creatorName"gorm:"type:varchar(110);not null"`
	CreatorId    int    `json:"creatorId"gorm:"not null"`
	JoinWay      string `json:"joinWay"`
	Date         string `json:"date"`
	Status       int    `json:"status"`
}
