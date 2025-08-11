package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetUserGradeList(c *gin.Context) {
	DB := common.GetDB()
	classesId := c.Query("classesId")
	examId := c.Query("examId")
	temp1, _ := strconv.Atoi(classesId)
	temp2, _ := strconv.Atoi(examId)
	if temp1 == 0 || temp2 == 0 {
		response.Fail(c, "查询学生考试情况失败", nil)
		return
	}
	var r = []model.UserTopic{}
	result := DB.Where("classes_id = ? AND exam_id = ?", temp1, temp2).Find(&r)
	response.Success(c, gin.H{"data": result}, "查询学生考试情况成功")

}
func GetStuExamInfo(c *gin.Context) {
	DB := common.GetDB()
	examId := c.Query("examId")
	classesId := c.Query("classesId")
	userId := c.Query("userId")
	temp1, _ := strconv.Atoi(examId)
	temp2, _ := strconv.Atoi(classesId)
	temp3, _ := strconv.Atoi(userId)
	println(temp1)
	println(temp2)
	println(temp3)
	var r = model.Exam{}
	result1 := DB.Where("exam_id = ?", temp1).Find(&r)
	var userexam = model.UserTopic{}
	result2 := DB.Where("exam_id = ? AND classes_id = ? AND user_id = ?", temp1, temp2, temp3).Find(&userexam)
	var examclass = model.ExamClasses{}
	result3 := DB.Where("exam_id = ? AND classes_id = ?", temp1, temp2).Find(&examclass)
	response.Success(c, gin.H{"edata": result1, "udata": result2, "ecdata": result3}, "查询试卷信息成功")
}
func TchMarkExam(c *gin.Context) {
	DB := common.GetDB()
	var usertopic = DTO.UserTopicDto{}
	c.ShouldBind(&usertopic)
	//answerTime := usertopic.AnswerTime
	classesId := usertopic.ClassesId
	//classesName := usertopic.ClassesName
	examId := usertopic.ExamId
	userId := usertopic.UserId
	println(classesId)
	println(examId)
	println(userId)
	temp1, _ := strconv.Atoi(classesId)
	temp2, _ := strconv.Atoi(examId)
	temp3, _ := strconv.Atoi(userId)
	grade := usertopic.Grade
	gradeAuto := usertopic.GradeAuto
	tY := time.Now().Year()
	tMo := time.Now().Month()
	tD := time.Now().Day()
	tH := time.Now().Hour()
	tMi := time.Now().Minute()
	tS := time.Now().Second()
	tNaS := time.Now().Nanosecond()
	curTimeDate := time.Date(tY, tMo, tD, tH, tMi, tS, tNaS, time.Local)
	fmt.Println("curTimeDate = ", curTimeDate)
	var user = model.UserTopic{}
	result := DB.Model(&user).Select("GradeAuto", "Grade", "MarkStatus").Where("exam_id = ? AND classes_id = ? AND user_id = ?", temp2, temp1, temp3).Updates(model.UserTopic{
		GradeAuto:  gradeAuto,
		Grade:      grade,
		MarkStatus: 1,
		MarkDate:   curTimeDate,
	})
	if result == nil {
		response.Fail(c, "审批考生试卷失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "提交考试信息成功")
}
