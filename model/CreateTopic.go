package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type CreateTopic struct {
	gorm.Model
	ExamId          int             `json:"examId"gorm:"not null;primaryKey"`
	ExamName        string          `json:"examName"gorm:"not null"`
	CreatorId       int             `json:"creatorId"gorm:"not null"`
	CreatorName     string          `json:"creatorName"gorm:"not null"`
	TopicNum        int             `json:"topicNum"gorm:"not null"`
	TopicTchDTOList json.RawMessage `json:"topicTchDTOList"gorm:"column:topicTchDTOList;type:json"`
}
