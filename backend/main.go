package main

import (
	"context"
	"fmt"
	"log"

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

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var ctx = context.TODO()
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

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to project gingo!",
		})
	})

	router.Run()
}
