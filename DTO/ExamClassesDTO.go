package DTO

import (
	"github.com/jinzhu/gorm"
)

type ExamClassesDto struct {
	gorm.Model
	ExamId         int `json:"examId"gorm:"not null"`
	ClassesId      int `json:"classesId"gorm:"not null"`
	ClassesArr     []int
	CLassesNameArr []string
	ExamName       string `json:"examName"gorm:"type:varchar(110);not null"`
	ClassesName    string `json:"classesName"gorm:"type:varchar(110);not null"`
	Deadline       string `json:"deadline"gorm:"type:varchar(110);not null"`
	StartDate      string `json:"startDate"gorm:"type:varchar(110)"`
	PublishAnswer  int    `json:"publishAnswer"gorm:"not null"`
	PublishScore   int    `json:"publishScore"gorm:"not null"`
	Status         int    `json:"status"gorm:"not null"`
	Releasing      int    `json:"releasing"gorm:"not null"`
	ExamTime       int    `json:"examTime"gorm:"not null"`
	TotalScore     int    `json:"totalScore"gorm:"not null"`
	PassMark       int    `json:"passMark"gorm:"not null"`
}
