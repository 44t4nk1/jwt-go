package main

import (
	"fmt"
	"log"
	"os"

	"github.com/44t4nk1/jwt-go/api/handler"
	"github.com/44t4nk1/jwt-go/api/middleware"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	claims, verified := extractClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjA0ODM0NzMsInVzZXJfaWQiOjF9.kI5G4U7q1FnmVAwfIkAIm3e0sUCZNmVgQoDnaHVIPBg")
	if verified {
		fmt.Println(claims)
	}
	router.GET("/home", middleware.IsAuthorised(handler.HomePage))
	router.POST("/login", handler.Login)
	log.Fatal(router.Run(":8080"))
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := os.Getenv("ACCESS_SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
