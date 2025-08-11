package DTO

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type UserTopicDTO struct {
	gorm.Model
	UserId        int
	UserName      string
	ClassesId     string
	ClassesName   string
	ExamId        string
	ExamName      string
	AnswerTime    int
	UserTopicList json.RawMessage
}
