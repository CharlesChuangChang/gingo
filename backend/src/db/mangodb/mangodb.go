package mangodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Client *mongo.Client

//存储在mongodb中的内容
type RecordLog struct {
	JobName   string    `bson:"jobName"`
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	Timepoint TimePoint `bson:"timepoint"`
}

type LogRecord struct {
	JobName string `bson:"jobName"`
}

type Person struct {
	Name string
	Age  int
}

func InsertRecord(client *mongo.Client, collect *mongo.Collection) (insertID primitive.ObjectID) {

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

func FindLog(collect *mongo.Collection) {
	// 创建需要过滤的条件
	logred := &LogRecord{
		JobName: "job1",
	}
	var skip int64 = 0  //从那个开始
	var limit int64 = 2 //炼制几个输出字段
	cursor, err := collect.Find(context.TODO(), logred, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		//创建需要反序列化成什么样子的结构体对象
		records := &RecordLog{}
		//反序列化
		err = cursor.Decode(records)
		if err != nil {
			fmt.Println(err)
			return
		}
		//打印
		fmt.Println(*records)
	}
}

func FindOne() {
	fmt.Println("FindOne:")
	col := Client.Database("my_db").Collection("my_col")
	var p Person

	//查询单条记录
	err := col.FindOne(context.TODO(), bson.D{{Key: "name", Value: "zhang3"}}).Decode(&p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)
}

func Find() {
	fmt.Println("Find:")
	opt := options.Find()
	opt.SetLimit(5)

	col := Client.Database("my_db").Collection("my_col")

	//Find查询多条记录，将会返回一个迭代器，使用完记得关闭
	cor, err := col.Find(context.TODO(), bson.D{{Key: "name", Value: "zhang3"}}, opt)
	if err != nil {
		log.Fatalln(err)
	}
	defer cor.Close(context.TODO())

	for cor.Next(context.TODO()) {
		var p Person
		cor.Decode(&p)
		fmt.Println(p)
	}
}
