package router

import (
	"DFES-Web/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	router.POST("/login", api.UserApiInstance.Login)
	//router.OPTIONS("/login", middleware.Cors)
	router.POST("/register", api.UserApiInstance.Register)
	//router.OPTIONS("/register", middleware.Cors)
}

func (ur *UserRouter) InitPrivateRouter(router *gin.RouterGroup) {
	router.GET("/userInfo", api.UserApiInstance.GetUserInfo)
}

var UserRouterInstance = new(UserRouter)
