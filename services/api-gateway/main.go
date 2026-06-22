package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.Use(gin.Recovery())

	public := router.Group("/api")
	{
		public.GET("/health", health)
	}

	private := router.Group("/api")
	{
		v1 := private.Group("/v1")
		v1.Use(AuthMiddleWare())
		{

		}
	}
	router.Run()
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func health(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
