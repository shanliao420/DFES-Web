package api

import (
	"DFES-Web/model/request"
	"DFES-Web/model/response"
	"DFES-Web/model/trans"
	"DFES-Web/service"
	"github.com/gin-gonic/gin"
	"log"
)

type UserApi struct {
}

func (ua *UserApi) Login(c *gin.Context) {
	var loginInfo request.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		log.Println("login bind json err:", err)
		response.FailWithMessage("请求参数有误", c)
		return
	}
	log.Println("login info:", loginInfo.Username)
	user := trans.LoginInfo2UserModel(&loginInfo)
	service.UserServiceInstance.Login(user, c)
}

func (ua *UserApi) Register(c *gin.Context) {
	var registerInfo request.RegisterInfo
	err := c.ShouldBindJSON(&registerInfo)
	if err != nil {
		log.Println("register bind json err:", err)
		response.FailWithMessage("请求参数有误", c)
		return
	}
	log.Println("register info:", registerInfo)

	user := trans.RegisterInfo2UserModel(&registerInfo)
	service.UserServiceInstance.Register(user, c)
	//log.Println("registry user[", registerInfo.Username, "] successful.")
}

var UserApiInstance = new(UserApi)
