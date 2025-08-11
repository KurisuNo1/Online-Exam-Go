package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type TopicList struct {
	gorm.Model
	CreatorId   int             `json:"creatorId"`
	CreatorName string          `json:"creatorName"`
	TopicData   json.RawMessage `json:"topicList"gorm:"column:topicList;type:json"`
}
