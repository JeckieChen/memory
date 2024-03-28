package service

import (
	"fmt"
	"memory/db"
)

type LoginInfo struct {
	UserName string
	Password string
}

func Login(loginInfo *LoginInfo) (tokenString string, err error) {
	username := loginInfo.UserName
	user := db.User{}
	db.DB.Where("username = ?", username).First(&user)
	fmt.Println(user)
	return "", nil

	// if user.Id == nil {
	// 	return "", nil
	// }
	// if user.Password == loginInfo.Password {
	// 	return util.GenerateJWT(username, 1000)
	// } else {
	// 	return "", nil
	// }

}
