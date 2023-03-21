package main

import (
	"context"
	"db/mangodbDriver"
	"fmt"
	"handler/loginHandler"
	"handler/registerHandler"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * @Author: cc
 * @Author: 346852628@qq.com
 * @Date: 2023/02/26
 * @Desc:backend project entry point
 */

func dbInit() {
	log.Println("dbInit")

}

func InitMangoDB() {
	log.Println("InitMangoDB")

	//创建上下文
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(10*time.Second, func() {
		cancelFunc()
	})

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	mangodbDriver.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = mangodbDriver.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

}

func main() {
	fmt.Println("welcome to project gingo!")

	InitMangoDB()

	collect := mangodbDriver.Client.Database("corn").Collection("jobs")
	//插入数据
	mangodbDriver.InsertRecord(mangodbDriver.Client, collect)
	//查询数据
	mangodbDriver.FindLog(collect)

	router := gin.Default()

	router.POST("/register", registerHandler.RegisterHandler)

	router.POST("/login", loginHandler.LoginHandler)
	authUrlRouter := router.Group("/gingo/V1")

	authUrlRouter.Use()

	router.Run("0.0.0.0:9091")
}
