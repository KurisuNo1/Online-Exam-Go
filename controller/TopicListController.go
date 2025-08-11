package controller

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateTopic(c *gin.Context) {
	DB := common.GetDB()
	var topic = model.CreateTopic{}
	c.ShouldBindJSON(&topic)
	creatorId := topic.CreatorId
	creatorName := topic.CreatorName
	examName := topic.ExamName
	topicNum := topic.TopicNum
	topicTchDTOList := topic.TopicTchDTOList
	println(creatorId)
	println(creatorName)
	println(topicNum)
	println(examName)
	DB.AutoMigrate(&topic)
	result := DB.Create(&model.CreateTopic{
		CreatorId:       creatorId,
		CreatorName:     creatorName,
		ExamName:        examName,
		TopicNum:        topicNum,
		TopicTchDTOList: topicTchDTOList,
	})
	if result == nil {
		response.Fail(c, "添加试题进入题库失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "添加试题进入题库成功")
}

func GetTopicList(c *gin.Context) {
	DB := common.GetDB()
	topicId := c.Query("keyword")
	temp, _ := strconv.Atoi(topicId)
	println(topicId)
	println(temp)
	if temp != 0 {
		var topic = model.CreateTopic{}
		result := DB.Where("exam_id = ?", temp).Find(&topic)
		if result == nil {
			response.Fail(c, "查询题库失败", nil)
			return
		}
		response.Success(c, gin.H{"data": result}, "查询题库成功")
		return
	}
	var topic = []model.CreateTopic{}
	result := DB.Find(&topic)
	if result == nil {
		response.Fail(c, "查询题库失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "查询题库成功")
}

func GetTopicListInformation(c *gin.Context) {
	DB := common.GetDB()
	examId := c.Query("examId")
	temp, _ := strconv.Atoi(examId)
	println(examId)
	println(temp)
	if temp == 0 {
		response.Fail(c, "查询题库信息失败", nil)
		return
	}
	var topic = model.CreateTopic{}
	result := DB.Where("exam_id = ?", temp).Find(&topic)
	response.Success(c, gin.H{"data": result}, "查询题库信息成功")
}
