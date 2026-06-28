package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunnybio/TaskForge/services/apigateway"
	"net/http"
)

func main() {

	router := gin.Default()
	router.Use(gin.Recovery())

	public := router.Group("/api")
	{
		public.GET("/health")
		public.POST("/login", login)
		public.POST("/signup", signup)
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
