package util

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DEFAULT_DATABASE_URI string
var DEFAULT_DATABASE_NAME string

var client *mongo.Client
var collection *mongo.Collection
var isConnected bool = false

// 初始化數據庫參數
func init() {
	DEFAULT_DATABASE_URI = os.Getenv("MONGO_DB_URI")
	DEFAULT_DATABASE_NAME = "emm-money-box"
}

// 開啓數據庫連綫
func OpenConnection(collectionName string) {

	// 檢查數據庫配置
	if DEFAULT_DATABASE_URI == "" {
		log.Fatal("Environment Value 'MONGO_DB_URI' not set.")
	} else {
		log.Output(0, "Connection established.")
		// log.Print("Using MONGO_DB_URI: " + DEFAULT_DATABASE_URI)
	}

	// 定義數據庫連綫
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DEFAULT_DATABASE_URI))
	if err != nil {
		log.Fatal(err)
	}

	// 設定數據集合
	collection = client.Database(DEFAULT_DATABASE_NAME).Collection(collectionName)
	isConnected = true
}

// 關閉數據庫連綫
func CloseConnection() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	isConnected = false
	log.Output(0, "Connection closed.")
}

// 檢查數據庫連綫
func checkConnection() {
	if !isConnected {
		log.Fatal("Empty Database Connection.")
	}
}

// 獲取單條數據
func QueryOne(filter bson.D) bson.M {

	checkConnection()

	var resultInBson bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&resultInBson)

	// 查詢失敗處理
	if err == mongo.ErrNoDocuments {
		fmt.Println("Record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return resultInBson
}

// 獲取多條數據
func QueryMany(filter bson.D) []bson.M {

	checkConnection()

	var resultInBsonArray []bson.M
	cursor, err := collection.Find(context.TODO(), filter)

	// 查詢失敗處理
	if err == mongo.ErrNoDocuments {
		fmt.Println("Records does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	if err := cursor.All(context.TODO(), &resultInBsonArray); err != nil {
		log.Fatal(err)
	}

	return resultInBsonArray
}
