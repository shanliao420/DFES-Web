package db

import (
	"github.com/go-redis/redis"
	"github.com/shanliao420/DFES-Go-Client/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	GlobalMySQLClient *gorm.DB
	GlobalRedisClient *redis.Client
	GlobalDFESClient  *client.DFESClient
)

func Init() {
	dsn := "root:333@tcp(127.0.0.1:3306)/dfes_web?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("error occur during mysql init:", err)
	}
	log.Println("mysql init successful.")
	GlobalMySQLClient = db

	cache := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   3,
	})
	log.Println("redis init successful.")
	GlobalRedisClient = cache

	DFESClient := client.NewDFESClient(":6000")
	GlobalDFESClient = DFESClient
}
