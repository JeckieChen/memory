package util

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("256-bit-secret")

type MyClaims struct {
	Subject string
	jwt.StandardClaims
}

func GenerateJWT(subject string, exp int) (string, error) {
	expireTime := time.Now().Add(time.Second * time.Duration(exp)).Unix()
	myClaims := &MyClaims{
		Subject: subject,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: expireTime,
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lym",
			NotBefore: 0,
			Subject:   "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenString, err := token.SignedString(secret)
	fmt.Println(tokenString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

func GetJwtInfo(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if myClaims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return myClaims, nil
	} else {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}
}
