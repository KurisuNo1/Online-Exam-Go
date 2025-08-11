package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	c.ShouldBind(&requestUser)
	//获取参数
	//name := requestUser.Name
	phone := requestUser.Phone
	password := requestUser.Password
	//数据验证
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "登陆成功")
}
func Register(c *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	c.ShouldBind(&requestUser)
	name := requestUser.Name
	email := requestUser.Email
	phone := requestUser.Phone
	password := requestUser.Password
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(email) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "邮箱不能为空")
		return
	}
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if isTelephoneExist(DB, phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该手机号已存在")
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
		Email:    email,
		Role:     "student",
	}
	DB.Create(&newUser)
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		return
	}
	response.Success(c, gin.H{"token": token}, "注册成功!")
	log.Println(name, phone, password, email)
	//response.Success(c, nil, "注册成功")
}

func isTelephoneExist(DB *gorm.DB, phone string) bool {
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	println()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"code": 200, "data": gin.H{"user": user}},
	})
}
func GetUserById(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的userId
	userId := claims.Id
	DB := common.GetDB()
	var user model.User
	result := DB.Where("id = ?", userId).First(&user)

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	} else {

	}
	response.Success(c, gin.H{"data": result}, "查询成功!")

}

func QueryUserByC_id(c *gin.Context) {
	DB := common.GetDB()
	classid := c.Query("c_id")
	temp, _ := strconv.Atoi(classid)
	var r = []model.UserClasses{}
	result := DB.Where("classes_id = ?", temp).Find(&r)
	if result == nil {
		response.Fail(c, "查询失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "按班级id查询成功")
}

func UpdateUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的classesId
	classesId := claims.Id
	DB := common.GetDB()
	var class model.Classes
	DB.Model(&class).Select("ClassName", "Introduction", "PeopleNum").Updates(model.Classes{ClassesName: "TEST", Introduction: "TEST", PeopleNum: 1})
	println(classesId)
	if class.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未查询到信息！"})
		c.Abort()
		return
	} else {
		response.Success(c, nil, "修改班级信息成功!")
	}
}

func GetRole(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的classesId
	userId := claims.Id
	DB := common.GetDB()
	var user model.User
	result := DB.Where("id = ?", userId).Select("role").First(&user)
	response.Success(c, gin.H{"data": result}, "获取用户角色成功！")
}

func ChangeRole(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, _ := common.ParseToken(tokenString)
	//验证通过后获取claim中的classesId
	userId := claims.Id
	status := claims.Status
	DB := common.GetDB()
	var user model.User
	println(status == "teacher")
	println(status == "student")
	if status == "teacher" {
		DB.Model(&user).Where("id = ?", userId).Update("role", "student")
		response.Success(c, nil, "修改用户角色为学生!")
		return
	} else if status == "student" {
		DB.Model(&user).Where("id = ?", userId).Update("role", "teacher")
		response.Success(c, nil, "修改用户角色为老师!")
		return
	}

}

func DeleteUser(c *gin.Context) {

}

func UpdateUserInfo(c *gin.Context) {
	DB := common.GetDB()
	var userInfo = DTO.UserDTO{}
	c.ShouldBindJSON(&userInfo)
	userId := userInfo.UserId
	//var id = int(userId)
	email := userInfo.Email
	name := userInfo.Name
	password := userInfo.Password
	phone := userInfo.Phone
	work := userInfo.Work
	sex := userInfo.Sex
	println(userId, email, name, password, phone, work, sex)
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	var u = model.User{}
	result := DB.Model(&u).Select("Name", "Email", "Password", "Phone", "Work", "Sex").Where("id = ?", userId).Updates(model.User{
		Name:     name,
		Email:    email,
		Password: string(hasedPassword),
		Phone:    phone,
		Work:     work,
		Sex:      sex,
	})
	if result == nil {
		response.Fail(c, "修改个人信息失败", nil)
		return
	}
	response.Success(c, gin.H{"data": result}, "修改个人信息成功")

}

// test
func UploadDeletephoto(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}
func Uploadphoto(c *gin.Context) {
	f, err := c.FormFile("uploadfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		dst := "C:/Go_workplace/awesomeProject/uploadfile/" + f.Filename
		c.SaveUploadedFile(f, dst)
		//c.JSON(http.StatusOK, gin.H{
		//	"message": fmt.Sprintln("'%s' uploaded!", f.Filename),
		//})
		response.Success(c, gin.H{"data": f.Filename}, "上传成功")
	}

}
