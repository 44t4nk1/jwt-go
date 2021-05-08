package handler

import (
	"net/http"

	"github.com/44t4nk1/jwt-go/api/middleware"
	"github.com/44t4nk1/jwt-go/api/models"
	"github.com/gin-gonic/gin"
)

var user = models.User{
	ID:       1,
	Username: "44t4nk1",
	Password: "PASSWORD",
}

//Login ...
func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON Provided")
		return
	}
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
