package mangodb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 存储在mongodb中的内容
type UserInfo struct {
	UserId    string    `bson:"UserId"`
	UserName  string    `bson:"username"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	HeadImage string    `bson:"image"`
	Timepoint TimePoint `bson:"timepoint"`
}

func AddNewUser(client *mongo.Client, c *gin.Context) error {
	body, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(body, &m)

	var userInfo UserInfo
	userInfo.UserId = "admin"
	userInfo.UserName = "admin"
	userInfo.Email = "admin"
	userInfo.Token = "admin"
	userInfo.HeadImage = "admin"
	userInfo.Timepoint = TimePoint{
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix(),
	}
	Collection := Client.Database("gingodb").Collection("users")
	_, err := Collection.InsertOne(context.TODO(), userInfo)
	if err != nil {
		return err
	}
	return nil
}

func userExist(collect *mongo.Collection, email string) (bool, error) {
	// 创建需要过滤的条件
	logred := &UserInfo{
		Email: email,
	}
	var skip int64 = 0  //从那个开始
	var limit int64 = 2 //炼制几个输出字段
	cursor, err := collect.Find(context.TODO(), logred, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		//创建需要反序列化成什么样子的结构体对象
		records := &RecordLog{}
		//反序列化
		err = cursor.Decode(records)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		//打印
		fmt.Println(*records)
	}
	return true, nil
}
func InsertUser(client *mongo.Client, collect *mongo.Collection) (insertID primitive.ObjectID) {

	collect = client.Database("corn").Collection("jobs")
	record := &RecordLog{
		JobName: "job1",
		Command: "main.go",
		Err:     "",
		Content: "Hello_World",
		Timepoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}
	insertRest, err := collect.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return
	}
	insertID = insertRest.InsertedID.(primitive.ObjectID)
	return insertID
}
