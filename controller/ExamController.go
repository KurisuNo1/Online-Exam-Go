package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTestPaperByU_id(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	//println(tokenString)
	_, claims, _ := common.ParseToken(tokenString)
	teacherid := claims.Id
	println(teacherid)
	if teacherid == 0 {
		response.Fail(c, "用户不存在", nil)
		return
	}
	DB := common.GetDB()
	var teach = []model.Exam{}
	result := DB.Where("creator_id = ?", teacherid).Find(&teach)
	response.Success(c, gin.H{"data": result}, "按教师id查询考试信息成功")
}

func GetExamByClasses(c *gin.Context) {
	DB := common.GetDB()
	classesId := c.Query("classesId")
	temp, _ := strconv.Atoi(classesId)
	println(temp)
	println(classesId)
	if temp == 0 {
		response.Fail(c, "查询失败", nil)
		return
	}
	var examclass = []model.ExamClasses{}
	result := DB.Where("classes_id = ?", temp).Find(&examclass)

	response.Success(c, gin.H{"data": result}, "按id查询班级成功")
}
func GetFinishExam(c *gin.Context) {
	//pagesize := c.Query("pageSize")
	//currentpage := c.Query("currentPage")
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	//println(tokenString)
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的userId
	userId := claims.Id
	if userId == 0 {
		response.Fail(c, "用户不存在!", nil)
		return
	}
	DB := common.GetDB()
	var exam []model.UserTopic
	//DB.AutoMigrate(exam)
	result := DB.Find(&exam, "user_id = ?", userId)
	response.Success(c, gin.H{"data": result}, "查询成绩成功！")

}
func GetTestPaperByTp_id(c *gin.Context) {
	DB := common.GetDB()
	examId := c.Query("examId")
	temp, _ := strconv.Atoi(examId)
	if temp == 0 {
		response.Fail(c, "按试卷id查询试卷信息失败", nil)
		return

	}
	var r = model.Exam{}
	result := DB.Where("exam_id = ?", temp).Find(&r)
	if result == nil {
		response.Fail(c, "按试卷id查询试卷信息失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "按试卷id查询试卷信息成功")
}

type Resu struct {
	ExamId        int    `json:"examId"gorm:"not null"`
	ClassesId     int    `json:"classesId"gorm:"not null"`
	ExamName      string `json:"examName"gorm:"type:varchar(110);not null"`
	ClassesName   string `json:"classesName"gorm:"type:varchar(110);not null"`
	Deadline      string `json:"deadline"gorm:"type:varchar(110);not null"`
	StartDate     string `json:"startDate"gorm:"type:varchar(110)"`
	PublishAnswer int    `json:"publishAnswer"gorm:"not null"`
	PublishScore  int    `json:"publishScore"gorm:"not null"`
	Status        int    `json:"status"gorm:"not null"`
	Releasing     string `json:"releasing"gorm:"type:varchar(110)"`

	Introduction string `json:"introduction"gorm:"type:varchar(110);not null"`
	PeopleNum    int    `json:"peopleNum"gorm:"not null"`
	CreatorName  string `json:"creatorName"gorm:"type:varchar(110);not null"`
	CreatorId    int    `json:"creatorId"gorm:"not null"`
	JoinWay      string `json:"joinWay"`
	Date         string `json:"date"`
}

func GetTestPaper(c *gin.Context) {
	temp := c.Query("userId")
	userId, _ := strconv.Atoi(temp)
	examId := c.Query("examId")
	classesId := c.Query("classesId")
	var e model.Exam
	var cl = model.ExamClasses{}
	var gr = []model.UserTopic{}
	DB := common.GetDB()
	u_result := DB.Where("user_id = ? AND exam_id = ? AND classes_id = ?", userId, examId, classesId).Find(&gr)
	e_result := DB.Where("exam_id = ? ", examId).Find(&e)
	c_result := DB.Where("exam_id = ? AND classes_id = ?", examId, classesId).Find(&cl)
	if e_result == nil || c_result == nil {
		response.Fail(c, "查询试卷信息失败", nil)
		return
	}
	response.Success(c, gin.H{"edata": e_result, "cdata": c_result, "udata": u_result}, "查询班级信息成功")
}
func CreateTestPaper(c *gin.Context) {
	DB := common.GetDB()
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的classesId
	creatorId := claims.Id
	println(creatorId)
	var testpaper = model.Exam{}
	c.ShouldBindJSON(&testpaper)
	var topicArr []*model.Topic
	//DB.AutoMigrate(&topicArr)
	//DB.AutoMigrate(&testpaper)
	Bytes, _ := testpaper.TopicTchDTOList.MarshalJSON()
	//DB.Create(&model.Exam{TopicTchDTOList: Bytes})
	json.Unmarshal(Bytes, &topicArr)
	for i := 0; i < len(topicArr); i++ {
		DB.Create(&model.Topic{
			Analysis:      topicArr[i].Analysis,
			Choice:        topicArr[i].Choice,
			CorrectAnswer: topicArr[i].CorrectAnswer,
			Difficulty:    topicArr[i].Difficulty,
			Question:      topicArr[i].Question,
			Required:      topicArr[i].Required,
			Score:         topicArr[i].Score,
			SubjectId:     topicArr[i].SubjectId,
			TopicType:     topicArr[i].TopicType,
		})
	}
	result := DB.Create(&model.Exam{
		AutoMack:        testpaper.AutoMack,
		CreatorName:     testpaper.CreatorName,
		DisruptOrder:    testpaper.DisruptOrder,
		ExamName:        testpaper.ExamName,
		PassMark:        testpaper.PassMark,
		PermitCopy:      testpaper.PermitCopy,
		RepeatTest:      testpaper.RepeatTest,
		SwitchPage:      testpaper.SwitchPage,
		Time:            testpaper.Time,
		TopicNum:        testpaper.TopicNum,
		TotalScore:      testpaper.TotalScore,
		CreatorId:       creatorId,
		TopicTchDTOList: Bytes,
	})
	response.Success(c, gin.H{"data": result}, "创建成功")
	return

}
func UpdateTestPaper(c *gin.Context) {
	DB := common.GetDB()
	var testinformation = model.Exam{}
	c.ShouldBindJSON(&testinformation)
	examid := testinformation.ExamId
	println(examid)
	result := DB.Model(&testinformation).Where("exam_id = ?", examid).Updates(model.Exam{
		ExamName:        testinformation.ExamName,
		CreatorId:       testinformation.CreatorId,
		CreatorName:     testinformation.CreatorName,
		Time:            testinformation.Time,
		SubjectId:       testinformation.SubjectId,
		SubjectName:     testinformation.SubjectName,
		TopicNum:        testinformation.TopicNum,
		TotalScore:      testinformation.TotalScore,
		PassMark:        testinformation.PassMark,
		PermitCopy:      testinformation.PermitCopy,
		SwitchPage:      testinformation.SwitchPage,
		DisruptOrder:    testinformation.DisruptOrder,
		RepeatTest:      testinformation.RepeatTest,
		AutoMack:        testinformation.AutoMack,
		Releasing:       testinformation.Releasing,
		TopicTchDTOList: testinformation.TopicTchDTOList,
	})
	if result == nil {
		response.Fail(c, "更新试卷信息失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "更新试卷信息成功")
}

func DeleteTestPaper(c *gin.Context) {
	DB := common.GetDB()
	testpaperId := c.Query("testPaperId")
	temp, _ := strconv.Atoi(testpaperId)
	if temp == 0 {
		response.Fail(c, "删除试卷失败", nil)
		return
	}
	result := DB.Where("exam_id = ?", temp).Delete(&model.Exam{})
	response.Success(c, gin.H{"data": result}, "删除试卷成功")

}
func SubmitTestPaper(c *gin.Context) {
	DB := common.GetDB()
	var usertopic = DTO.UserTopicDTO{}
	c.ShouldBind(&usertopic)
	answerTime := usertopic.AnswerTime
	classesId := usertopic.ClassesId
	classesName := usertopic.ClassesName
	examId := usertopic.ExamId
	userId := usertopic.UserId
	userName := usertopic.UserName
	userTopicList := usertopic.UserTopicList
	temp1, _ := strconv.Atoi(classesId)
	temp2, _ := strconv.Atoi(examId)
	println(temp1)
	println(temp2)
	DB.AutoMigrate(&model.UserTopic{})
	//var ut = model.UserTopic{}
	r := model.UserTopic{AnswerTime: answerTime, ClassesId: temp1, ClassesName: classesName, ExamId: temp2, UserId: userId, UserName: userName, UserTopicList: userTopicList, ExamStatus: 1}
	result := DB.Select("AnswerTime", "ClassesId", "ClassesName", "ExamId", "UserId", "UserName", "UserTopicList", "ExamStatus", " CreatedAt", "UpdatedAt").Create(&r)
	response.Success(c, gin.H{"data": result}, "提交考试信息成功")
}
