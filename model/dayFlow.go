package model

import (
	"log"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayFlowEntity struct {
	year      string
	month     string
	day       string
	cashFlows []string
}

func GetDayFlowEntity(date time.Time) DayFlowEntity {

	var entity DayFlowEntity

	filter := bson.D{
		primitive.E{Key: "year", Value: date.Year()},
		primitive.E{Key: "month", Value: date.Format("01")},
		primitive.E{Key: "day", Value: date.Day()},
	}
	util.OpenConnection("dayFlow")

	// fixme: 如何對queryResult判空
	queryResult := util.QueryOne(filter)

	err := mapstructure.Decode(queryResult, &entity)
	if err != nil {
		log.Fatal(err)
	}

	util.CloseConnection()
	return entity
}
