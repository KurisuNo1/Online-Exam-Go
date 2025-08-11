package model

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	TopicId       int    `json:"topicId"gorm:"not null"`
	CreatorId     int    `json:"creatorId"gorm:"not null"`
	SubjectId     string `json:"subjectId"gorm:"not null"`
	SubjectName   string `json:"subjectName"gorm:"type:varchar(110);not null"`
	Question      string `json:"question"gorm:"type:varchar(110);not null"`
	Choice        string `json:"choice"gorm:"type:varchar(110);not null"`
	Photo         string `json:"photo"gorm:"type:varchar(110);not null"`
	TopicType     int    `json:"topicType"gorm:"not null"`
	CorrectAnswer string `json:"correctAnswer"gorm:"type:varchar(110);not null"`
	Score         int    `json:"score"gorm:"not null"`
	Difficulty    string `json:"difficulty"gorm:"type:varchar(110);not null"`
	Analysis      string `json:"analysis"gorm:"type:varchar(110);not null"`
	Required      int    `json:"required"gorm:"not null"`
}
