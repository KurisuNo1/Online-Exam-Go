package DTO

import "github.com/jinzhu/gorm"

type TopicTchDTOList struct {
	gorm.Model
	Analysis      string `json:"analysis"gorm:"type:varchar(110);not null"`
	Choice        string `json:"choice"gorm:"type:varchar(110);not null"`
	CorrectAnswer string `json:"correctAnswer"gorm:"type:varchar(110);not null"`
	Difficulty    string `json:"difficulty"gorm:"type:varchar(110);not null"`
	Question      string `json:"question"gorm:"type:varchar(110);not null"`
	Required      int    `json:"required"gorm:"type:varchar(110);not null"`
	Score         int    `json:"score"gorm:"type:varchar(110);not null"`
	SubjectId     string `json:"subjectId"gorm:"type:varchar(110);not null"`
	TopicType     int    `json:"topicType"gorm:"type:varchar(110);not null"`
}
