package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Exam struct {
	gorm.Model
	ExamId          int             `json:"examId"gorm:"not null;primaryKey"`
	ExamName        string          `json:"examName"gorm:"not null"`
	CreatorId       int             `json:"creatorId"`
	CreatorName     string          `json:"creatorName"gorm:"not null"`
	Time            int             `json:"time"gorm:"not null"`
	SubjectId       int             `json:"subjectId"gorm:"not null"`
	SubjectName     string          `json:"subjectName"gorm:"not null"`
	TopicNum        int             `json:"topicNum"gorm:"not null"`
	TotalScore      int             `json:"totalScore"gorm:"not null"`
	PassMark        int             `json:"passMark"gorm:"not null"`
	PermitCopy      int             `json:"permitCopy"gorm:"not null"`
	SwitchPage      int             `json:"switchPage"gorm:"not null"`
	DisruptOrder    int             `json:"disruptOrder"gorm:"not null"`
	RepeatTest      string          `json:"repeatTest"gorm:"not null"`
	AutoMack        int             `json:"autoMack"gorm:"not null"`
	Releasing       int             `json:"releasing"gorm:"not null"`
	TopicTchDTOList json.RawMessage `json:"topicTchDTOList"gorm:"column:topicTchDTOList;type:json"`
}
