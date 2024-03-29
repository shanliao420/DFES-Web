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
	engin.Use(middleware.TokenCheck)

	engin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "successful",
		})
	})

	anyGroup := engin.Group("/api")
	anyGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "true",
		})
	})

	router.UserRouterInstance.InitUserRouter(anyGroup)

	err := engin.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
