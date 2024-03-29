package middleware

import (
	"DFES-Web/model/response"
	"DFES-Web/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func TokenCheck(c *gin.Context) {
	authInfo := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authInfo, "Basic ")
	log.Println("auth by token:", token)
	if !utils.ExistsToken(token) {
		response.FailWithMessage("用户未登陆", c)
		return
	}
	userInfo := utils.GetToken(token)
	c.Set("user-info", userInfo)
	c.Next()
}
