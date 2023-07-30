package database

import (
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"log"
	"sync"
)

var once = sync.Once{}
var defaultDatabaseUri string
var defaultDatabaseName string
var isConnected = false

var CashFlowTableName = "cash_flow"
var CategoryTableName = "category"

func initMongoDbConnection() {
	defaultDatabaseUri = util.GetConfigByKey("db.mongodb.url")
	defaultDatabaseName = util.GetConfigByKey("db.name")
}

func initMySqlConnection() {
	defaultDatabaseUri = util.GetConfigByKey("db.mysql.url")
	defaultDatabaseName = util.GetConfigByKey("db.name")
}

func checkDbConnection() {
	if !isConnected {
		log.Fatal("empty database connection.")
	}
}
