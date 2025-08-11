package model

import "github.com/jinzhu/gorm"

type Approval struct {
	gorm.Model
	UserId       int    `json:"userId"gorm:"not null"`
	UserName     string `json:"userName"gorm:"type:varchar(110)"`
	TeacherId    int    `json:"teacherId"gorm:"default:'1'"`
	ApprovalRole string `json:"approvalRole"gorm:"type:varchar(110)"`
	ClassesId    int    `json:"classesId"gorm:"not null"`
	ClassesName  string `json:"classesName"`
	JoinWay      string `json:"joinWay"`
	Status       int    `json:"status"gorm:"default:null"`
}
