package util

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	DEFAULT_DATABASE_URI = GetConfigByKey("db.url")
	DEFAULT_DATABASE_NAME = GetConfigByKey("db.name")
}

// 開啓數據庫連綫
func OpenMongoDbConnection(collectionName string) {

	// 檢查數據庫配置
	if DEFAULT_DATABASE_URI == "" {
		log.Fatal("Environment Value 'MONGO_DB_URI' not set.")
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

// 關閉數據庫連綫 FIXME: 存在空指針問題
func CloseMongoDbConnection() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	isConnected = false
	// log.Output(0, "Connection closed.")
}

// 檢查數據庫連綫
func checkMongoDbConnection() {
	if !isConnected {
		log.Fatal("Empty Database Connection.")
	}
}

func GetOneInMongoDb(filter bson.D) bson.M {

	checkMongoDbConnection()

	var resultInBson bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&resultInBson)

	// 查詢失敗處理
	if err == mongo.ErrNoDocuments {
		Logger.Warnln("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return resultInBson
}

func GetManyInMongoDb(filter bson.D) []bson.M {

	checkMongoDbConnection()

	var resultInBsonArray []bson.M
	cursor, err := collection.Find(context.TODO(), filter)

	// 查詢失敗處理
	if err == mongo.ErrNoDocuments {
		Logger.Warnln("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	if err2 := cursor.All(context.TODO(), &resultInBsonArray); err2 != nil {
		log.Fatal(err2)
	}

	return resultInBsonArray
}

func InsertOneInMongoDb(data bson.D) primitive.ObjectID {

	checkMongoDbConnection()

	/* result:
	 *	type InsertOneResult struct {
	 *		InsertedID primitive.ObjectID
	 *	}
	 */
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}

	return result.InsertedID.(primitive.ObjectID)
}

// Deprecated: use UpdateMany() instead.
func UpdateOneInMongoDb(filter, data bson.D) bool {

	checkMongoDbConnection()

	updateData := bson.D{
		primitive.E{Key: "$set", Value: data},
	}

	// If the filter matches multiple documents, one will be selected from the matched set and MatchedCount will equal 1.
	result, err := collection.UpdateOne(context.TODO(), filter, updateData)
	if err != nil {
		panic(err)
	}

	return result.ModifiedCount == 1
}

func UpdateManyInMongoDb(filter, data bson.D) int64 {

	checkMongoDbConnection()

	updateData := bson.D{
		primitive.E{Key: "$set", Value: data},
	}
	// Upsert disable by default.
	result, err := collection.UpdateMany(context.TODO(), filter, updateData)
	if err != nil {
		panic(err)
	}

	return result.ModifiedCount
}

// Deprecated: use DeleteMany() instead.
func DeleteOneInMongoDb(filter bson.D) bool {

	checkMongoDbConnection()

	// If the filter matches multiple documents, one will be selected from the matched set.
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return result.DeletedCount == 1
}

func DeleteManyInMongoDb(filter bson.D) int64 {

	checkMongoDbConnection()

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return result.DeletedCount
}
