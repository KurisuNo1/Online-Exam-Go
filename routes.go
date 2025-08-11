package main

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/say_hello", controller.Test)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/topicList", controller.TopicList1)
	//ApprovalController
	r.GET("/approval/list", controller.ApprovalList)
	r.POST("/approval", controller.Approval)
	r.GET("/approval", controller.Approval)
	//TopicListController
	r.POST("/createTopic", controller.CreateTopic)
	r.GET("/getTopicList", controller.GetTopicList)
	r.GET("/getTopicInfo", controller.GetTopicListInformation)
	//ClassesController
	r.GET("/queryClasses", controller.QueryClasses)
	r.GET("/queryClasses1", controller.QueryClasses1)
	r.GET("/queryClassesList", controller.QueryClassesList)
	r.GET("/queryClassesList1", controller.QueryClassesList1)
	r.GET("/queryClassesList2", controller.QueryClassesList2)
	r.POST("/joinClasses", controller.JoinClasses)
	r.POST("/outClasses", controller.OutClasses)
	r.DELETE("/outClassesByTeacher", controller.OutClassesByTeacher)
	r.POST("/createClasses", controller.CreateClasses)
	r.PUT("/updateClasses", controller.UpdateClasses)
	r.POST("/classes/fuzzyQuery", controller.ClassesFuzzyQuery)
	r.DELETE("/deleteClasses", controller.DeleteClasses)
	//ExamController
	r.GET("/getTestPaperByU_id", controller.GetTestPaperByU_id)
	r.GET("/getExamByClasses", controller.GetExamByClasses)
	r.GET("/getFinishExam", controller.GetFinishExam)
	r.GET("/getTestPaperByTp_id", controller.GetTestPaperByTp_id)
	r.GET("/getTestPaper", controller.GetTestPaper)
	r.POST("/createTestPaper", middleware.AuthMiddleware(), controller.CreateTestPaper)
	r.POST("/updateTestPaper", controller.UpdateTestPaper)
	r.POST("/submitTestPaper", controller.SubmitTestPaper)
	r.DELETE("/deleteTestPaper", controller.DeleteTestPaper)
	//MarkExamController
	r.GET("/getUserGradeList", controller.GetUserGradeList)
	r.GET("/getStuExamInfo", controller.GetStuExamInfo)
	r.PUT("/tchMarkExam", controller.TchMarkExam)
	//ReleaseExamController
	r.GET("/getReleaseInfo", controller.GetReleaseInfo)
	r.POST("/releaseTest", controller.ReleaseTest)
	r.POST("/updateReleaseTest", controller.UpdateReleaseTest)
	r.DELETE("/cancelRelease", controller.CancelRelease)
	r.PUT("/updateReleasedTest", controller.UpdateReleasedTest)
	//UploadController
	r.POST("/upload", controller.Uploadphoto)
	r.POST("/upload/delete", controller.UploadDeletephoto)
	//UserController
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/getUserById", middleware.AuthMiddleware(), controller.GetUserById)
	r.GET("/getUserById", middleware.AuthMiddleware(), controller.GetUserById)
	r.GET("/queryUserByC_id", controller.QueryUserByC_id)
	r.POST("/updateUser", middleware.AuthMiddleware(), controller.UpdateUser)
	r.POST("/getRole", middleware.AuthMiddleware(), controller.GetRole)
	r.PUT("/changeRole", middleware.AuthMiddleware(), controller.ChangeRole)
	r.POST("/deleteUser", middleware.AuthMiddleware(), controller.DeleteUser)
	r.PUT("/updateUserInfo", controller.UpdateUserInfo)
	return r
}
