package auth

import (
	
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var jwtkey=[]byte("golang")

type Claims struct{
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(username,role string)(string ,error)  {
	expireTime:=time.Now().Add(1*time.Hour)

	claim:=&Claims{
		Username: username,
		Role :role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString(jwtkey)
	
}

func ValidateToken(tokenString string)(*Claims, error){
	claims:=&Claims{}

	token,err:=jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token) (interface{}, error) {
		return jwtkey,nil
	})
	if err!=nil || !token.Valid{
		return nil,errors.New("invalid or expired token")
	}
	return claims,nil
}