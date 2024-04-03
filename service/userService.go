package service

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"DFES-Web/model/response"
	"DFES-Web/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
}

func (us *UserService) Register(user *do.UserModel, c *gin.Context) {
	cnt := int64(0)
	db.GlobalMySQLClient.Where("username = ?", user.Username).First(&do.UserModel{}).Count(&cnt)
	if cnt > 0 {
		response.FailWithMessage("注册失败，用户名已存在", c)
		return
	}
	db.GlobalMySQLClient.Where("phone = ?", user.Phone).First(&do.UserModel{}).Count(&cnt)
	if cnt > 0 {
		response.FailWithMessage("注册失败，预留手机号已存在", c)
		return
	}
	user.Password = utils.BcryptHash(user.Password)
	err := db.GlobalMySQLClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := FileSystemServiceInstance.InitUserRootNode(user); err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("register err occur in db insert:", err)
		response.FailWithMessage("注册失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("注册成功😊", c)
}

func (us *UserService) Login(user *do.UserModel, c *gin.Context) {
	var userInDB do.UserModel
	if err := db.GlobalMySQLClient.Where("username = ?", user.Username).First(&userInDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("用户名或密码不正确", c)
			return
		}
		log.Println("query user in db process err:", err)
		response.FailWithMessage("登陆失败，请稍后再试", c)
		return
	}
	if !utils.BcryptCheck(user.Password, userInDB.Password) {
		response.FailWithMessage("用户名或密码不正确", c)
		return
	}
	token := utils.SetOrGetToken(&userInDB)
	c.Header("Authorization", "Basic "+token)
	response.OkWithMessage("登陆成功", c)
}

var UserServiceInstance = new(UserService)
