package model

import (
	"strconv"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayFlowEntity struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	CashFlows []primitive.ObjectID `json:"cashFlows" bson:"cashFlows"`
	Day       int                  `json:"day" bson:"day"`
	Month     int                  `json:"month" bson:"month"`
	Year      int                  `json:"year" bson:"year"`
}

func GetDayFlowByDate(date time.Time) DayFlowEntity {

	var entity DayFlowEntity

	// Golang有趣的日期转换机制，不过转出来是String，数据库是int，所以转一下类型。
	monthInInt, _ := strconv.Atoi(date.Format("01"))

	// 查询条件为指定的年月日，设计上一天只会对应一笔dayFlow数据。
	filter := bson.D{
		primitive.E{Key: "year", Value: date.Year()},
		primitive.E{Key: "month", Value: monthInInt},
		primitive.E{Key: "day", Value: date.Day()},
	}

	// 打开dayFlow的数据表连线
	util.OpenConnection("dayFlow")

	// 获取查询结果并转入结构对象
	queryResult := util.QueryOne(filter)
	bsonBytes, _ := bson.Marshal(queryResult)
	bson.Unmarshal(bsonBytes, &entity)

	// fixme: 数据库关闭连线好像有点问题
	// util.CloseConnection()

	return entity
}
