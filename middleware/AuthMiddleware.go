package middleware

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够",
			})
			c.Abort()
			return
		}
		//验证通过后获取claim中的userId
		userId := claims.Id
		userName := claims.UserName
		status := claims.Status

		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		//println(userId)
		if len(userName) != 0 {
			//println(userName)

		}
		if len(status) != 0 {
			//println(status)
		}

		c.Set("user", user)
		c.Next()
	}
}
