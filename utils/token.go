package utils

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultExpireTime = 5 * time.Hour
	UserTokenPrefix   = "user-token-"
	TokenUserPrefix   = "token-user-"
)

type TokenModel struct {
	Token string
	do.UserModel
}

func GenerateToken() string {
	uid, _ := uuid.NewUUID()
	token := strings.ToUpper(uid.String())
	token = strings.ReplaceAll(token, "-", "")
	return token
}

func SetOrGetToken(user *do.UserModel) string {
	tokenUserKey := UserTokenPrefix + strconv.FormatUint(user.ID, 10)
	ret := db.GlobalRedisClient.Exists(tokenUserKey)
	if ret.Val() != 0 {
		tokenVal := db.GlobalRedisClient.Get(tokenUserKey)
		var tokenModel TokenModel
		_ = jsoniter.UnmarshalFromString(tokenVal.Val(), &tokenModel)
		userTokenKey := TokenUserPrefix + tokenModel.Token
		db.GlobalRedisClient.Expire(tokenUserKey, DefaultExpireTime)
		db.GlobalRedisClient.Expire(userTokenKey, DefaultExpireTime)
		return tokenModel.Token
	}
	token := GenerateToken()
	userTokenKey := TokenUserPrefix + token
	tokenModel := &TokenModel{
		Token:     token,
		UserModel: *user,
	}
	json, _ := jsoniter.MarshalToString(tokenModel)
	db.GlobalRedisClient.Set(tokenUserKey, json, DefaultExpireTime)
	db.GlobalRedisClient.Set(userTokenKey, token, DefaultExpireTime)
	return token
}

func GetToken(token string) *do.UserModel {
	ret := db.GlobalRedisClient.Get(TokenUserPrefix + token)
	if ret.Err() != nil || ret.Val() == "" {
		return nil
	}
	tRet := db.GlobalRedisClient.Get(TokenUserPrefix + ret.Val())
	var tokenModel TokenModel
	_ = jsoniter.UnmarshalFromString(tRet.Val(), &tokenModel)
	return &(tokenModel.UserModel)
}

func ExistsToken(token string) bool {
	return db.GlobalRedisClient.Exists(TokenUserPrefix+token).Val() != int64(0)
}
