package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func PrintURI(c *gin.Context) {
	uri := c.Request.RequestURI
	log.Println("uri:", uri)
	c.Next()
}
