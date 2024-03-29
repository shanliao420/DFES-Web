package utils

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"errors"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strings"
	"time"
)

const (
	DefaultExpireTime = 5 * time.Hour
)

func GenerateToken() string {
	uid, _ := uuid.NewUUID()
	token := strings.ToUpper(uid.String())
	token = strings.ReplaceAll(token, "-", "")
	return token
}

func SetToken(token string, user *do.UserModel) {
	json, err := jsoniter.MarshalToString(user)
	if err != nil {
		log.Println("set toke on marshal err:", err)
		return
	}
	db.GlobalRedisClient.Set(token, json, DefaultExpireTime)
}

func GetToken(token string) *do.UserModel {
	ret := db.GlobalRedisClient.Get(token)
	if errors.Is(ret.Err(), redis.Nil) {
		return nil
	}
	var user do.UserModel
	err := jsoniter.UnmarshalFromString(ret.String(), &user)
	if err != nil {
		log.Println("get token unmarshal err:", err)
		return nil
	}
	return &user
}

func ExistsToken(token string) bool {
	return db.GlobalRedisClient.Exists(token).Val() != int64(0)
}
