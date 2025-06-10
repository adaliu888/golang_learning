package service

import (
	DB "golang_learning/mynewpro/db"
	"golang_learning/mynewpro/jwt"
	"golang_learning/mynewpro/middlewave"
	"golang_learning/mynewpro/pojo"

	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//var userList = []pojo.User{}

// get user
func FindAllUsers(c *gin.Context) {
	//c.JSON(200, userList)
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

// get user by id
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "error: user not found")
	}
	log.Printf("%+v", user)
	c.JSON(http.StatusOK, user)

}

// post user
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	//添加到数据库

	//userList = append(userList, user)
	newuser := pojo.PostUser(user)

	c.JSON(http.StatusOK, newuser)
	//return
}

// delete user
func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "error")
	}
	c.JSON(http.StatusOK, "deleted successfully")

}

/* delete user
userId, err := strconv.Atoi(c.Param("id"))
if err != nil {
	c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
	return
}
for _, user := range userList {
	log.Printf("%+v", user)

	userList = append(userList[:userId], userList[userId+1:]...)
	c.JSON(http.StatusOK, "Successfully deleted user")
	return

}*/

// update user
func UpdateUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "error")

	}

	c.JSON(http.StatusOK, user)

}

// login user
func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := pojo.CheckUserPassword(name, password)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "error")
		return
	}

	userId := strconv.Itoa(user.Id)

	// 保留会话认证（向后兼容）
	middlewave.SaveSession(c, userId)

	// 生成JWT令牌
	token, err := jwt.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// 生成刷新令牌
	refreshToken, err := jwt.GenerateRefreshToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "login successfully",
		"user":          user,
		"session":       middlewave.GetSession(c), // 向后兼容
		"token":         token,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
	})
}

// logout user
func Logout(c *gin.Context) {
	// 会话注销 (向后兼容)
	middlewave.ClearSession(c)

	// JWT 令牌注销
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString := authHeader[7:]
		// 尝试撤销令牌
		claims, err := jwt.ValidateToken(tokenString)
		if err == nil {
			jwt.RevokeToken(tokenString)
			jwt.ClearUserRefreshTokens(claims.UserID)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logout successfully",
	})
}

func CheckUserSession(c *gin.Context) {
	sessionID := middlewave.GetSession(c)
	if sessionID == "0" {
		c.JSON(http.StatusUnauthorized, "error")

		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "check session successfully",
		"user":    middlewave.GetSession(c),
	})

}

//func Logout(c *gin.Context) {

/*
	beforUser := pojo.User{}
	err := c.BindJSON(&beforUser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	userId, _ := strconv.Atoi(c.Param("id"))

	for key, user := range userList {
		if userId == user.Id {
			userList[key] = beforUser
			log.Printf("%+v", userList[key])
			c.JSON(http.StatusOK, "Successfully updated user")
			return
		}
	}
	c.JSON(http.StatusNotFound, "error: user not updated")
}
*/
//redis user

func RedisOneUser(c *gin.Context) {
	id := c.Param("id")
	if id == "0" {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	user := pojo.User{}
	DB.DBConnect.Find(&user, id)
	c.Set("dbResult", user)

}
