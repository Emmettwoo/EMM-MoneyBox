package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"time"
)

type CategoryEntity struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	ParentId   primitive.ObjectID `json:"parent_id" bson:"parent_id"`
	Name       string             `json:"name" bson:"name"`
	Remark     string             `json:"remark" bson:"remark"`
	CreateTime time.Time          `json:"create_time" bson:"create_time"`
	ModifyTime time.Time          `json:"modify_time" bson:"modify_time"`
}

func (entity CategoryEntity) IsEmpty() bool {
	return reflect.DeepEqual(entity, CategoryEntity{})
}

func (entity CategoryEntity) ToString() string {

	return "[ " +
		"Id: " + entity.Id.Hex() +
		", Name: " + entity.Name +
		" ]"
}
