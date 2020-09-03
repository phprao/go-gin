package jwtUtil

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/**
DEMO

token := jwtUtil.CreateToken(120, "15919906312")
fmt.Println(token)
info, _ := jwtUtil.ParseToken(token)
fmt.Println(info.UserId, info.Phone, info.ExpiresAt)
 */

const DEFAULT_EXPITE = 3600

type CustomClaims struct {
	UserId int
	Phone  string
	jwt.StandardClaims
}

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("please set JWT_SECRET in env")
	}
}

func CreateToken(userId int, phone string) string {
	secret := os.Getenv("JWT_SECRET")

	exp := os.Getenv("JWT_EXPIRE")
	expire, err := strconv.Atoi(exp)
	if err != nil || expire == 0 {
		expire = DEFAULT_EXPITE
	}
	customClaims := &CustomClaims{
		UserId: userId, //用户id
		Phone:  phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(),
			Issuer:    "",
		},
	}

	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	fmt.Printf("token: %v\n", tokenString)

	return tokenString
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || token == nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
