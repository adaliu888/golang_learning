package service

import (
	DB "golang_learning/mynewpro/db"
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
	//userId := strconv.Itoa(user.Id)
	userId := strconv.Itoa(user.Id)

	middlewave.SaveSession(c, userId)
	//token, err := jwt.GenerateToken(userId)
	//if err != nil {
	//c.JSON(500, gin.H{"error": "failed to generate token"})
	//return
	//}
	// 存储Token到Session
	// 可以在这里设置Cookie或者使用其他方式存储Token
	// 例如，使用SetCookie设置Token到Cookie
	// c.SetCookie("auth_token", token, 3600, "/", "yourdomain.com", false, true)
	//c.SetCookie("auth_token", token, 3600, "/", "127.0.0.1/8080", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
		"user":    user,
		"session": middlewave.GetSession(c), //use session
		//"token": token,
	})
	// 登录成功，重定向到目标页面
	// 例如，重定向到用户的个人主页 "/user/profile"
	c.Redirect(http.StatusFound, "/user/profile")

	// 注意：使用c.Redirect后不需要再调用c.JSON或其他返回响应的方法

}

// logout user
func Logout(c *gin.Context) {
	middlewave.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successfully",
		"session": middlewave.GetSession(c),
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
