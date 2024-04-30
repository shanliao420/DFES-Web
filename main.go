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
	engin.Use(middleware.Cors)

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
	router.FileSystemRouterInstance.InitFileSystemRouter(privateGroup)

	err := engin.Run(":9080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
