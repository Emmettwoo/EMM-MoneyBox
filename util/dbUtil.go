package util

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DEFAULT_DATABASE_URL string = "mongodb://emmettwoo:emmettwoo@cluster0-shard-00-00.it4fc.mongodb.net:27017,cluster0-shard-00-01.it4fc.mongodb.net:27017,cluster0-shard-00-02.it4fc.mongodb.net:27017/hello-vercel?replicaSet=atlas-13kbxi-shard-0&ssl=true&authSource=admin"
const DEFAULT_DATABASE_NAME string = "emm-money-box"

var client *mongo.Client
var collection *mongo.Collection
var isConnected bool = false

// 開啓數據庫連綫
func OpenConnection(collectionName string) {

	// 定義數據庫連綫
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DEFAULT_DATABASE_URL))
	if err != nil {
		log.Fatal(err)
	}

	// 設定數據集合
	collection = client.Database(DEFAULT_DATABASE_NAME).Collection(collectionName)
	isConnected = true
}

// 關閉數據庫連綫
func CloseConnection() {
	log.Output(0, "Connection closed.")
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	isConnected = false
}

// 檢查數據庫連綫
func checkConnection() {
	if !isConnected {
		log.Fatal("Empty Database Connection.")
	}
}

// 獲取單條數據
func QueryOne(filter bson.D) bson.M {

	fmt.Println(filter)
	checkConnection()

	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	// 查詢失敗處理
	if err == mongo.ErrNoDocuments {
		fmt.Println("Record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return result
}

// 獲取多條數據
func QueryMany(filter bson.D) []bson.M {

	checkConnection()

	var results []bson.M
	cursor, err := collection.Find(context.TODO(), filter)

	// 查詢失敗處理
	if err == mongo.ErrNoDocuments {
		fmt.Println("Record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results
}
