package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//User ...
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//A sample use
var user = User{
	ID:       1,
	Username: "44t4nk1",
	Password: "PASSWORD",
}

//Login ...
func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON Provided")
	}
}

func main() {
	router.POST("/login", Login)
	log.Fatal(router.Run(":8080"))
}
