package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type UserTopic struct {
	gorm.Model
	UserId        int             `json:"userId"`
	UserName      string          `json:"userName"`
	ClassesId     int             `json:"classesId"`
	ClassesName   string          `json:"classesName"`
	ExamId        int             `json:"examId"`
	TopicId       int             `json:"topicId"`
	UserAnswer    string          `json:"userAnswer"`
	UserScore     int             `json:"userScore"`
	TopicStatus   int             `json:"topicStatus"`
	AnswerTime    int             `json:"answerTime"`
	Grade         int             `json:"grade"`
	GradeAuto     int             `json:"gradeAuto"`
	MarkStatus    int             `json:"markStatus"`
	MarkDate      time.Time       `json:"markDate"gorm:"default:null"`
	ExamStatus    int             `json:"examStatus"`
	AnswerDate    string          `json:"answerDate"`
	UserTopicList json.RawMessage `json:"userTopicList"gorm:"column:userTopicList;type:json"`
}
