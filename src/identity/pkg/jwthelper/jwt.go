package jwthelper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



func GenerateJwt(username string)(string,error){
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":username,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	})

	
	tokenString,err:= token.SignedString([]byte(os.Getenv("TokenKey")))

	return tokenString,err
}


func VerifyToken(tokenStr string) (bool,error){
	token,err := jwt.Parse(tokenStr, func(t *jwt.Token)(interface{},error){
		return []byte(os.Getenv("TokenKey")),nil
	})

	if err!=nil{
		return false,err
	}

	return token.Valid,nil
}
