package database

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func OpenMongoDbConnection(collectionName string) {

	// check and init database setting
	once.Do(initMongoDbConnection)
	if defaultDatabaseUri == "" {
		log.Fatal("environment value 'MONGO_DB_URI' not set")
	}

	// 定義數據庫連綫
	var err error
	client, err = mongo.Connect(context.TODO(),
		options.Client().ApplyURI(defaultDatabaseUri).SetMaxPoolSize(50))
	if err != nil {
		log.Fatal(err)
	}

	// 設定數據集合
	collection = client.Database(defaultDatabaseName).Collection(collectionName)
	isConnected = true
	util.Logger.Debugln("database connection created")
}

// CloseMongoDbConnection fixme: take action after CRUD, ensure every action have open a new connection.
func CloseMongoDbConnection() {

	// do nothing if not connected
	if !isConnected || reflect.DeepEqual(client, mongo.Client{}) {
		isConnected = false
		return
	}
	// close the connection
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	isConnected = false
	util.Logger.Debugln("database connection closed")
}

func GetOneInMongoDB(filter bson.D) bson.M {

	checkDbConnection()

	var resultInBson bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&resultInBson)

	// 查詢失敗處理
	if errors.Is(err, mongo.ErrNoDocuments) {
		// Logger.Warnln("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return resultInBson
}

func GetManyInMongoDB(filter bson.D) []bson.M {

	checkDbConnection()

	var resultInBsonArray []bson.M
	cursor, err := collection.Find(context.TODO(), filter)

	// 查詢失敗處理
	if errors.Is(err, mongo.ErrNoDocuments) {
		// Logger.Warnln("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	if err2 := cursor.All(context.TODO(), &resultInBsonArray); err2 != nil {
		log.Fatal(err2)
	}

	return resultInBsonArray
}

func CountInMongoDB(filter bson.D) int64 {

	checkDbConnection()

	result, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return result
}

func InsertOneInMongoDB(data bson.D) primitive.ObjectID {

	checkDbConnection()

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

func UpdateManyInMongoDB(filter, data bson.D) int64 {

	checkDbConnection()

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

func DeleteManyInMongoDB(filter bson.D) int64 {

	checkDbConnection()

	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return result.DeletedCount
}
