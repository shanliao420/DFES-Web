package middleware

import (
	"DFES-Web/model/response"
	"DFES-Web/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

const (
	UserInfoCacheKey = "user-info"
)

func TokenCheck(c *gin.Context) {
	authInfo := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authInfo, "Basic ")
	log.Println("auth by token:", token)
	if !utils.ExistsToken(token) {
		response.FailWithMessage("用户未登陆", c)
		c.Abort()
		return
	}
	userInfo := utils.GetToken(token)
	c.Set(UserInfoCacheKey, userInfo)
	c.Next()
}
