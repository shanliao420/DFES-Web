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
	userTokenKey := UserTokenPrefix + strconv.FormatUint(user.ID, 10)
	ret := db.GlobalRedisClient.Exists(userTokenKey)
	if ret.Val() != 0 {
		tokenVal := db.GlobalRedisClient.Get(userTokenKey)
		var tokenModel TokenModel
		_ = jsoniter.UnmarshalFromString(tokenVal.Val(), &tokenModel)
		tokenUserKey := TokenUserPrefix + tokenModel.Token
		db.GlobalRedisClient.Expire(userTokenKey, DefaultExpireTime)
		db.GlobalRedisClient.Expire(tokenUserKey, DefaultExpireTime)
		return tokenModel.Token
	}
	token := GenerateToken()
	tokenUserKey := TokenUserPrefix + token
	tokenModel := &TokenModel{
		Token:     token,
		UserModel: *user,
	}
	json, _ := jsoniter.MarshalToString(tokenModel)
	db.GlobalRedisClient.Set(tokenUserKey, strconv.FormatUint(user.ID, 10), DefaultExpireTime)
	db.GlobalRedisClient.Set(userTokenKey, json, DefaultExpireTime)
	return token
}

func GetToken(token string) *do.UserModel {
	ret := db.GlobalRedisClient.Get(TokenUserPrefix + token)
	if ret.Err() != nil || ret.Val() == "" {
		return nil
	}
	tRet := db.GlobalRedisClient.Get(UserTokenPrefix + ret.Val())
	var tokenModel TokenModel
	_ = jsoniter.UnmarshalFromString(tRet.Val(), &tokenModel)
	return &(tokenModel.UserModel)
}

func ExistsToken(token string) bool {
	return db.GlobalRedisClient.Exists(TokenUserPrefix+token).Val() != int64(0)
}
