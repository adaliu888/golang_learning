package pojo

import (
	"golang_learning/mynewpro/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"id" `
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func FindAllUsers() []User {
	var users []User
	db.DBConnect.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User
	db.DBConnect.Where("id = ?", userId).First(&user)
	return user
}

// add user
func PostUser(user User) User {
	db.DBConnect.Create(&user)
	return user
}

// delete user
func DeleteUser(userId string) User {
	user := User{}
	db.DBConnect.Where("id = ?", userId).Delete(&User{})
	return user
}

// update user
func UpdateUser(userId string, user User) User {
	db.DBConnect.Model(&user).Where("id = ?", userId).Update("user", user)
	return user
}

func CheckUserPassword(username string, password string) User {
	var user User
	db.DBConnect.Where("username = ? AND password = ?", username, password).First(&user)
	return user
}
