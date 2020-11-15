package handler

import (
	"net/http"

	"github.com/44t4nk1/jwt-go/api/models"
	"github.com/gin-gonic/gin"
)

//HomePage ...
func HomePage(c *gin.Context) {
	var Response = models.Response{
		Error:   false,
		Message: "Home Page",
	}
	c.JSON(http.StatusOK, Response)
}
