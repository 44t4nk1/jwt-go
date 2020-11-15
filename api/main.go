package main

import (
	"log"

	"github.com/44t4nk1/jwt-go/api/handler"
	"github.com/44t4nk1/jwt-go/api/middleware"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/home", middleware.IsAuthorised(handler.HomePage))
	router.POST("/login", handler.Login)
	log.Fatal(router.Run(":8080"))
}
