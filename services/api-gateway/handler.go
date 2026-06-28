package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Signup(c *gin.Context) {

}
func Health(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
