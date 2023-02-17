package main

import (
	"context"
	"fmt"
	"gingo/src/db/mangodb"
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

func main() {
	fmt.Println("welcome to project gingo!")

	//创建上下文
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(10*time.Second, func() {
		cancelFunc()
	})

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	defer client.Disconnect(ctx)

	collect := client.Database("corn").Collection("jobs")
	//插入数据
	mangodb.InsertRecord(client, collect)
	//查询数据
	mangodb.FindLog(collect)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to project gingo!",
		})
	})

	router.Run()
}
