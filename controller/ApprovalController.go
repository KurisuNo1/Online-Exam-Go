package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Test(c *gin.Context) {
	//输出json结果给调用方
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func ApprovalList(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的userId
	Id := claims.Id
	status := claims.Status
	if Id == 0 {
		response.Fail(c, "用户不存在!", nil)
		return
	}
	if status == "teacher" {
		DB := common.GetDB()
		var teacherapproval []model.Approval
		DB.AutoMigrate(&teacherapproval)
		result := DB.Find(&teacherapproval, "teacher_id = ?", Id)
		response.Success(c, gin.H{"data": result}, "查询学生申请成功！")
		return
	}
	if status == "student" {
		DB := common.GetDB()
		var userapproval []model.Approval
		DB.AutoMigrate(&userapproval)
		result := DB.Find(&userapproval, "user_id = ?", Id)
		response.Success(c, gin.H{"data": result}, "查询申请成功！")
		return
	}

}

func Approval(c *gin.Context) {
	DB := common.GetDB()
	var approva = DTO.ApprovalDto{}
	c.ShouldBindJSON(&approva)
	approvalRole := approva.ApprovalRole
	classesId := approva.ClassesId
	classesName := approva.ClassesName
	userId := approva.UserId
	userName := approva.UserName
	status := approva.Status
	//println(approvalRole)
	//println(classesId)
	//println(classesName)
	//println(userId)
	//println(userName)
	//println(status)
	if userId == 0 || status == 0 {
		response.Success(c, nil, "申请失败")
		return
	}
	var r = model.Approval{}
	DB.AutoMigrate(&r)
	result := DB.Model(&r).Where("user_id = ? AND classes_id = ?", userId, classesId).Updates(map[string]interface{}{"status": status})
	if result == nil {
		response.Fail(c, "审批失败", nil)
		return
	}
	if status == 1 {
		var userclass = model.UserClasses{}
		var class = model.Classes{}
		DB.AutoMigrate(&userclass)
		result1 := DB.Create(&model.UserClasses{
			ClassesId:   classesId,
			ClassesName: classesName,
			UserId:      userId,
			UserName:    userName,
			Position:    approvalRole,
		})
		if result1 == nil {
			response.Fail(c, "审批失败", nil)
			return
		}
		//q2 := DB.Table("classes").Select("people_num").Where("classes_id = ?", classesId).SubQuery()
		//DB.AutoMigrate(&class)
		//DB.Model(&model.Classes{}).Where("classes_id = ?", classesId).Updates(&model.Classes{PeopleNum: DB.Table("classes").Select("people_num").Where("classes_id = ?", classesId).SubQuery()})
		DB.Model(&class).Where("classes_id = ?", classesId).UpdateColumn("people_num", gorm.Expr("people_num + ?", 1))
	}
	response.Success(c, gin.H{"data": result}, "审批成功")
}

func TopicList1(c *gin.Context) {
	DB := common.GetDB()
	var td = model.TopicList{}
	c.ShouldBindJSON(&td)
	topicList := td.TopicData
	creatorId := td.CreatorId
	creatorName := td.CreatorName
	DB.AutoMigrate(&td)
	result := DB.Create(&model.TopicList{
		CreatorId:   creatorId,
		CreatorName: creatorName,
		TopicData:   topicList,
	})
	if result == nil {
		response.Fail(c, "录入题库失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "录入题库成功")

}
