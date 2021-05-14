package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/44t4nk1/jwt-go/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySigningKey = []byte(os.Getenv("ACCESS_SECRET"))

func IsAuthorised(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if c.GetHeader("Token") != "" {
			token, err := jwt.Parse(c.GetHeader("Token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an Error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				var Response = models.Response{
					Error:   true,
					Message: "Invalid Signature",
				}
				c.JSON(http.StatusUnauthorized, Response)
			}
			if token.Valid {
				endpoint(c)
			}
		} else {
			var Response = models.Response{
				Error:   true,
				Message: "No token provided",
			}
			c.JSON(http.StatusUnauthorized, Response)
		}
	})
}
