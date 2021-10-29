package model

import (
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CashFlowEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Amount   float64            `json:"amount" bson:"amount"`
	Category string             `json:"category" bson:"category"`
	Desc     string             `json:"desc" bson:"desc"`
	Remark   string             `json:"remark" bson:"remark"`
}

func GetCashFlowByObjectId(objectId primitive.ObjectID) CashFlowEntity {

	var entity CashFlowEntity

	// 查询条件为指定的cashFlow id，设计上一天只会对应一笔dayFlow数据。
	filter := bson.D{
		primitive.E{Key: "_id", Value: objectId},
	}

	// 打开cashFlow的数据表连线
	util.OpenConnection("cashFlow")

	// 获取查询结果并转入结构对象
	queryResult := util.QueryOne(filter)
	bsonBytes, _ := bson.Marshal(queryResult)
	bson.Unmarshal(bsonBytes, &entity)

	// fixme: 数据库关闭连线好像有点问题
	// util.CloseConnection()

	return entity
}

func GetCashFlowsByObjectIdArray(objectIdArray []primitive.ObjectID) []CashFlowEntity {

	var entity CashFlowEntity
	var entityArray []CashFlowEntity

	filter := bson.D{
		primitive.E{Key: "_id", Value: bson.M{"$in": objectIdArray}},
	}

	// 打开cashFlow的数据表连线
	util.OpenConnection("cashFlow")

	// 获取查询结果并转入结构对象
	queryResultArray := util.QueryMany(filter)
	for _, queryResult := range queryResultArray {
		bsonBytes, _ := bson.Marshal(queryResult)
		bson.Unmarshal(bsonBytes, &entity)
		entityArray = append(entityArray, entity)
	}

	// fixme: 数据库关闭连线好像有点问题
	// util.CloseConnection()

	return entityArray

}
