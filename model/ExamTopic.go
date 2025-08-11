package model

import "github.com/jinzhu/gorm"

type ExamTopic struct {
	gorm.Model
	ExamId    int `json:"examId"gorm:"not null"`
	TopicId   int `json:"topicId"gorm:"not null"`
	TopicType int `json:"topicType"gorm:"not null"`
}
