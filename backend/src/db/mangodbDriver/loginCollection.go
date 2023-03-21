package mangodbDriver

import (
	"context"
	"fmt"
	loginModel "myModel/dbModel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginCollectionGetUserInfo(client *mongo.Client, user *loginModel.LoginInfo) (*loginModel.UserInfo, error) {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"email": user.Email}
	singleDoc := Collection.FindOne(context.TODO(), filter)
	//var data map[string]interface{}
	userData := &loginModel.UserInfo{}
	err := singleDoc.Decode(&userData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(userData)

	return userData, nil
}

func LogomCollectionCheckPassword(client *mongo.Client, user *loginModel.LoginInfo) (bool, error) {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"email": user.Email}
	singleDoc := Collection.FindOne(context.TODO(), filter)
	var data map[string]interface{}

	err := singleDoc.Decode(&data)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(data)
	return true, nil
}
