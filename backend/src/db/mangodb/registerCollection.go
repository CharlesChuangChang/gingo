package mangodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterCollectionAdd(client *mongo.Client, user *UserInfo) error {
	fmt.Println("RegisterCollectionAdd")
	Collection := client.Database("gingodb").Collection("users")
	_, err := Collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RegisterCollectionFindByName(client *mongo.Client, user *UserInfo) error {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"username": user.UserName}
	singleDoc := Collection.FindOne(context.TODO(), filter)
	var data map[string]interface{}
	err := singleDoc.Decode(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(data)
	return nil
}

func RegisterCollectionFindById(client *mongo.Client, user *UserInfo) error {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"userid": user.UserId}
	singleDoc := Collection.FindOne(context.TODO(), filter)
	var data map[string]interface{}
	err := singleDoc.Decode(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(data)
	return nil
}

func RegisterCollectionFindByEmail(client *mongo.Client, user *UserInfo) error {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"email": user.Email}
	singleDoc := Collection.FindOne(context.TODO(), filter)
	var data map[string]interface{}
	err := singleDoc.Decode(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(data)
	return nil
}

func RegisterCollectionDeleteById(client *mongo.Client, user *UserInfo) error {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"userid": user.UserId}
	_, err := Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func RegisterCollectionDeleteByName(client *mongo.Client, user *UserInfo) error {
	Collection := client.Database("gingodb").Collection("users")
	filter := bson.M{"username": user.UserName}
	_, err := Collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func RegisterCollectionUpdate(client *mongo.Client, user *UserInfo) error {
	Collection := client.Database("gingodb").Collection("users")
	selector := bson.M{"userid": user.UserId}
	_, err := Collection.UpdateOne(context.TODO(), selector, user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func userExist(collect *mongo.Collection, email string) (bool, error) {
	return true, nil
}
