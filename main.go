package main

import (
	"DFES-Web/db"
	"DFES-Web/middleware"
	"DFES-Web/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engin := gin.Default()
	db.Init()

	engin.Use(middleware.PrintURI)

	engin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "successful",
		})
	})

	rootGroup := engin.Group("/api")
	rootGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "true",
		})
	})

	router.UserRouterInstance.InitUserRouter(rootGroup)

	privateGroup := rootGroup.Group("/private")
	privateGroup.Use(middleware.TokenCheck)
	router.UserRouterInstance.InitPrivateRouter(privateGroup)

	err := engin.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
