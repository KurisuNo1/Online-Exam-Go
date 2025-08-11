package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetReleaseInfo(c *gin.Context) {
	DB := common.GetDB()
	classesId := c.Query("classesId")
	examId := c.Query("examId")
	temp1, _ := strconv.Atoi(classesId)
	temp2, _ := strconv.Atoi(examId)
	if temp1 == 0 || temp2 == 0 {
		response.Fail(c, "获得考试信息失败", nil)
		return
	}
	var r = model.ExamClasses{}
	result := DB.Where("classes_id = ? AND exam_id = ?", temp1, temp2).Find(&r)
	response.Success(c, gin.H{"data": result}, "获得考试信息成功")

}
func ReleaseTest(c *gin.Context) {
	DB := common.GetDB()
	var updateinfo = DTO.ExamClassesDto{}
	c.ShouldBind(&updateinfo)
	examId := updateinfo.ExamId
	examTime := updateinfo.ExamTime
	totalScore := updateinfo.TotalScore
	examName := updateinfo.ExamName
	releasing := updateinfo.Releasing
	var classesArr []int = updateinfo.ClassesArr
	var classesNameArr []string = updateinfo.CLassesNameArr
	startDate := updateinfo.StartDate
	deadline := updateinfo.Deadline
	publishScore := updateinfo.PublishScore
	publishAnswer := updateinfo.PublishAnswer
	passmark := updateinfo.PassMark
	println(len(classesArr))
	for i := 0; i < len(classesArr); i++ {
		temp1 := classesArr[i]
		temp2 := classesNameArr[i]
		println(temp1)
		println(temp2)
		DB.AutoMigrate(&model.ExamClasses{})
		DB.Select("ExamId", "ClassesId", "ClassesName", "StartDate", "Deadline", "PublishScore", "PublishAnswer", "Releasing", "ExamTime", "ExamName", "TotalScore", "PassMark").Create(&model.ExamClasses{
			ExamId:        examId,
			ClassesId:     temp1,
			ClassesName:   temp2,
			StartDate:     startDate,
			Deadline:      deadline,
			PublishScore:  publishScore,
			PublishAnswer: publishAnswer,
			Releasing:     releasing,
			ExamTime:      examTime,
			ExamName:      examName,
			TotalScore:    totalScore,
			PassMark:      passmark})
	}
	response.Success(c, nil, "发布考试成功！")

}
func UpdateReleaseTest(c *gin.Context) {
	DB := common.GetDB()
	var updateinfo = DTO.ExamClassesDto{}
	c.ShouldBind(&updateinfo)
	examId := updateinfo.ExamId
	var classesArr []int = updateinfo.ClassesArr
	startDate := updateinfo.StartDate
	deadline := updateinfo.Deadline
	publishScore := updateinfo.PublishScore
	publishAnswer := updateinfo.PublishAnswer
	for i := 0; i < len(classesArr); i++ {
		println(classesArr[i])
		m := map[string]interface{}{
			"StartDate":     startDate,
			"Deadline":      deadline,
			"PublishScore":  publishScore,
			"PublishAnswer": publishAnswer,
		}
		DB.Model(&model.ExamClasses{}).Select("StartDate", "Deadline", "PublishScore", "PublishAnswer").Where("exam_id = ? AND classes_id = ?", examId, classesArr[i]).Updates(m)
	}
	response.Success(c, nil, "发布考试成功！")
}
func UpdateReleasedTest(c *gin.Context) {
	DB := common.GetDB()
	var updateinfo = model.ExamClasses{}
	c.ShouldBind(&updateinfo)
	examId := updateinfo.ExamId
	classesId := updateinfo.ClassesId
	startDate := updateinfo.StartDate
	deadline := updateinfo.Deadline
	publishScore := updateinfo.PublishScore
	publishAnswer := updateinfo.PublishAnswer
	m := map[string]interface{}{
		"StartDate":     startDate,
		"Deadline":      deadline,
		"PublishScore":  publishScore,
		"PublishAnswer": publishAnswer,
	}
	result := DB.Model(&model.ExamClasses{}).Select("StartDate", "Deadline", "PublishScore", "PublishAnswer").Where("exam_id = ? AND classes_id = ?", examId, classesId).Updates(m)
	if result == nil {
		response.Fail(c, "更新考试信息失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "更新考试信息成功")
}
func CancelRelease(c *gin.Context) {
	DB := common.GetDB()
	testpaperid := c.Query("tp_id")
	classesid := c.Query("c_id")
	temp1, _ := strconv.Atoi(testpaperid)
	temp2, _ := strconv.Atoi(classesid)
	if temp1 == 0 || temp2 == 0 {
		response.Fail(c, "取消发布考试失败", nil)
		return
	}
	var cancel = model.ExamClasses{}
	result := DB.Where("exam_id = ? AND classes_id = ?", temp1, temp2).Unscoped().Delete(&cancel)
	response.Success(c, gin.H{"data": result}, "取消发布考试成功")
}
