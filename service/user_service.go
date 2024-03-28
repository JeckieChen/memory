package service

import (
	"fmt"
	"memory/db"
	"memory/util"
)

type LoginInfo struct {
	UserName string
	Password string
}

func Login(loginInfo *LoginInfo) (tokenString string, err error) {
	username := loginInfo.UserName
	fmt.Println(username)
	user := db.User{}
	db.DB.Where("username = ?", username).First(&user)
	if user.Password == loginInfo.Password {
		token, _ := util.GenerateJWT(username, 300)
		return token, nil
	}
	fmt.Println(user)
	return "", err
}

type UserInfo struct {
	UserName string
	Password string
}

func Register(userInfo *UserInfo) (status bool, err error) {
	user := db.User{Username: userInfo.UserName, Password: userInfo.Password}
	result := db.DB.Create(user)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		return true, nil
	}
	return false, result.Error
}
