package model

import (
	"github.com/jinzhu/gorm"
)

type UserGrade struct {
	gorm.Model
	UserId      int    `json:"userId"`
	ClassesId   int    `json:"classesId"`
	ExamId      int    `json:"examId"`
	UserName    string `json:"userName"`
	ClassesName string `json:"classesName"`
	ExamName    string `json:"examName"`
	Grade       int    `json:"grade"`
	GradeAuto   string `json:"gradeAuto"`
	AnswerDate  string `json:"answerDate"`
	AnswerTime  int    `json:"answerTime"`
	MarkStatus  int    `json:"markStatus"`
	ExamStatus  int    `json:"examStatus"`
	MarkDate    string `json:"markDate"`
}
