package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func QueryClasses1(c *gin.Context) {
	DB := common.GetDB()
	result := DB.Find(&model.Classes{})
	if result == nil {
		response.Fail(c, "查询班级失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "查询班级成功")
}
func QueryClasses(c *gin.Context) {
	classesId := c.Query("classesId")
	temp, _ := strconv.Atoi(classesId)
	DB := common.GetDB()
	var class = model.Classes{}
	result := DB.Where("classes_id = ?", temp).Find(&class)
	if result == nil {
		response.Fail(c, "查询班级信息失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "查询班级信息成功")
}

func QueryClassesList(c *gin.Context) {
	classid := c.Query("keyword")
	userId := c.Query("userId")
	userRole := c.Query("userRole")
	temp1, _ := strconv.Atoi(classid)
	temp2, _ := strconv.Atoi(userId)
	if temp1 == 0 && temp2 == 0 && userRole == "" {
		DB := common.GetDB()
		var info = []model.Classes{}
		result := DB.Find(&info)
		response.Success(c, gin.H{"data": result}, "查询班级成功！")
		return
	}
	if temp1 != 0 {
		DB := common.GetDB()
		var info = []model.Classes{}
		result := DB.Where("classes_id = ?", temp1).Find(&info)
		if result == nil {
			response.Fail(c, "查询班级失败1", nil)
			return
		}
		response.Success(c, gin.H{"data": result}, "查询班级成功！")
		return
	}
	if userRole == "teacher" {
		DB := common.GetDB()
		var teacherclass = []model.Classes{}
		result := DB.Where("creator_id = ?", temp2).Find(&teacherclass)
		if result == nil {
			response.Fail(c, "查询班级失败2", nil)
			return
		}
		response.Success(c, gin.H{"data": result}, "查询班级成功！")
		return
	}
	if userRole == "student" {
		DB := common.GetDB()
		var stuclasses = []model.Classes{}

		//q2 := DB.Table("user_classes").Select("classes_id").Where("user_id = ?", temp2).SubQuery()
		//result := DB.Where("classes_id = ?", q2).Find(&stuclasses)
		result := DB.Where("classes_id IN ?", DB.Table("user_classes").Select("classes_id").Where("user_id = ?", temp2).SubQuery()).Find(&stuclasses)
		if result == nil {
			response.Fail(c, "查询班级失败3", nil)
			return
		}
		response.Success(c, gin.H{"data": result}, "查询班级成功！")
		return
	}

}
func QueryClassesList1(c *gin.Context) {
	DB := common.GetDB()
	examid := c.Query("examId")
	temp, _ := strconv.Atoi(examid)
	if temp == 0 {
		response.Fail(c, "查询失败！", nil)
		return
	}
	var info = []model.ExamClasses{}
	result := DB.Where("exam_id = ?", temp).Find(&info)
	response.Success(c, gin.H{"data": result}, "查询成功！")

}
func QueryClassesList2(c *gin.Context) {
	DB := common.GetDB()
	var info = []model.ExamClasses{}
	result := DB.Find(&info)
	response.Success(c, gin.H{"data": result}, "查询成功！")

}
func JoinClasses(c *gin.Context) {
	DB := common.GetDB()
	var app = DTO.UserClassesDTO{}
	c.ShouldBindJSON(&app)
	classesId := app.ClassesId
	temp, _ := strconv.Atoi(classesId)
	userId := app.UserId
	userName := app.UserName
	approvalRole := app.ApprovalRole
	joinWay := app.JoinWay
	classesName := app.ClassesName
	peopleNum := app.PeopleNum + 1
	println(userId)
	println(approvalRole)
	println(classesId)
	println(temp)
	println(joinWay)
	println(classesName)
	println(peopleNum)
	//temp1, _ := strconv.Atoi(classid)
	//temp2, _ := strconv.Atoi(userId)
	if temp == 0 || userId == 0 {
		response.Fail(c, "申请加入班级失败", nil)
		return
	}
	if joinWay == "all" {
		r := model.UserClasses{ClassesId: temp, UserId: userId, UserName: userName, Position: approvalRole, ClassesName: classesName}
		DB.AutoMigrate(&model.UserClasses{})
		var updateNum = model.Classes{}
		//DB.AutoMigrate(&updateNum)
		DB.Model(&updateNum).Where("classes_id = ?", temp).Updates(model.Classes{PeopleNum: peopleNum})
		result := DB.Create(&r)
		if result == nil {
			response.Fail(c, "加入班级失败", nil)
			return
		}
		response.Success(c, gin.H{"data": result}, "加入班级成功")
		return
	}
	if joinWay == "apply" {
		r := model.Approval{ClassesId: temp, UserId: userId, UserName: userName, ApprovalRole: approvalRole, JoinWay: joinWay, ClassesName: classesName}
		DB.AutoMigrate(&model.Approval{})
		result := DB.Create(&r)
		if result == nil {
			response.Fail(c, "申请加入班级失败", nil)
			return
		}
		response.Success(c, gin.H{"data": result}, "申请加入班级成功")
	}

}

func OutClasses(c *gin.Context) {
	var uc = DTO.ClassesUserDTO{}
	c.ShouldBindJSON(&uc)
	userId := uc.UserId
	classesId := uc.ClassesId
	//println(userId)
	//println(classesId)
	DB := common.GetDB()
	//var id = model.Classes{ClassesId: temp}
	result := DB.Unscoped().Where("classes_id = ? AND user_id = ?", classesId, userId).Delete(&model.UserClasses{})
	if result == nil {
		response.Fail(c, "退出班级失败！", nil)
		return
	}
	if classesId != 0 {
		//println(classesId)
		var class = model.Classes{}
		//DB.AutoMigrate(&class)
		DB.Model(&class).Where("classes_id = ?", classesId).UpdateColumn("people_num", gorm.Expr("people_num - ?", 1))
		response.Success(c, nil, "退出班级成功")
	}

}

func OutClassesByTeacher(c *gin.Context) {
	DB := common.GetDB()
	userId := c.Query("u_id")
	classesId := c.Query("c_id")
	temp1, _ := strconv.Atoi(userId)
	temp2, _ := strconv.Atoi(classesId)
	var r = model.UserClasses{}
	result := DB.Where("user_id = ? AND classes_id = ?", temp1, temp2).Unscoped().Delete(&r)
	if result == nil {
		response.Fail(c, "按班级id和学生id踢出学生失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "按班级id和学生id踢出学生成功")
}

func CreateClasses(c *gin.Context) {
	DB := common.GetDB()
	creatorId := c.Query("creatorId")
	temp, _ := strconv.Atoi(creatorId)
	classesName := c.Query("classesName")
	creatorName := c.Query("creatorName")
	joinWay := c.Query("joinway")
	introduction := c.Query("introduction")
	var users = model.Classes{ClassesName: classesName, CreatorName: creatorName, CreatorId: temp, JoinWay: joinWay, Introduction: introduction}
	result := DB.Create(&users)
	if result == nil {
		response.Fail(c, "创建班级失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "创建班级成功")
}

func UpdateClasses(c *gin.Context) {
	DB := common.GetDB()
	var updateClasses = model.Classes{}
	c.ShouldBindJSON(&updateClasses)
	classesId := updateClasses.ClassesId
	classesName := updateClasses.ClassesName
	introduction := updateClasses.Introduction
	joinWay := updateClasses.JoinWay
	println(classesId)
	println(classesName)
	println(introduction)
	println(joinWay)
	result := DB.Model(&updateClasses).Where("classes_id = ?", classesId).Updates(model.Classes{ClassesName: classesName, Introduction: introduction, JoinWay: joinWay, ClassesId: classesId})
	if classesId == 0 {
		response.Fail(c, "更新失败！", nil)
		return
	} else {
		response.Success(c, gin.H{"data": result}, "更新成功")
		return
	}

}

func ClassesFuzzyQuery(c *gin.Context) {

}

func DeleteClasses(c *gin.Context) {
	var classid = c.Query("c_id")
	temp, _ := strconv.Atoi(classid)
	DB := common.GetDB()
	//var id = model.Classes{ClassesId: temp}
	result := DB.Unscoped().Where("classes_id = ?", temp).Delete(&model.Classes{})
	if result == nil {
		response.Fail(c, "删除班级失败！", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "删除班级成功")
}
