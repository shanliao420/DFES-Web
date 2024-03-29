package router

import (
	"DFES-Web/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur UserRouter) InitUserRouter(router *gin.RouterGroup) {
	router.POST("/login", api.UserApiInstance.Login)
	router.POST("/register", api.UserApiInstance.Register)
}

var UserRouterInstance = new(UserRouter)
