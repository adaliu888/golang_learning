package service

import (
	"log"
	"mynewpro/middlewave"
	"mynewpro/pojo"
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
	middlewave.SaveSession(c, userId)
	c.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
		"user":    user,
		"session": middlewave.GetSession(c),
	})
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
