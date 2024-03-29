package trans

import (
	"DFES-Web/model/do"
	"DFES-Web/model/request"
)

func RegisterInfo2UserModel(registerInfo *request.RegisterInfo) *do.UserModel {
	return &do.UserModel{
		Username: registerInfo.Username,
		Password: registerInfo.Password,
		Email:    registerInfo.Email,
		Phone:    registerInfo.Phone,
	}
}

func LoginInfo2UserModel(loginInfo *request.LoginInfo) *do.UserModel {
	return &do.UserModel{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
	}
}
