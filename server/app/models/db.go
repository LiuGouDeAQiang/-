package model

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

// 数据库操作都放在这里
var MB *mongo.Client
var Conn *gorm.DB
var Rdb *redis.Client

func NewMongoDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://192.168.198.128:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	MB = client

	fmt.Println("Connected to MongoDB!")
}
func NewMysql() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "127.0.0.1:3306", "esports")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	Conn = conn
}

func NewRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.198.128:6379",
		Password: "",
		DB:       0,
	})
	Rdb = rdb
	//初始化session
	Store, _ = redisstore.NewRedisStore(context.TODO(), Rdb)
	return
}

//用户 --> 客户端应用程序: 输入手机号和验证码
//客户端应用程序 --> 服务器: 发送手机号和验证码
//服务器 --> 客户端应用程序: 验证手机号和验证码
//客户端应用程序 --> 用户: 登录成功
// 邮箱验证码部分

