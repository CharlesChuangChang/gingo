package main

import (
	"context"
	"fmt"
	"gingo/src/db/mangodb"
	"gingo/src/handler/registerHandler"
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
	mangodb.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = mangodb.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	defer mangodb.Client.Disconnect(ctx)
}

func main() {
	fmt.Println("welcome to project gingo!")

	InitMangoDB()

	collect := mangodb.Client.Database("corn").Collection("jobs")
	//插入数据
	mangodb.InsertRecord(mangodb.Client, collect)
	//查询数据
	mangodb.FindLog(collect)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to project gingo!",
		})
	})

	router.POST("/register", registerHandler.RegisterHandler)

	router.GET("/login", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Login success!",
		})
	})

	router.Run("0.0.0.0:7000")
}
