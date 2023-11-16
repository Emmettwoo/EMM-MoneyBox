package util

import "os"

var configurationMap map[string]string

func init() {
	configurationMap = make(map[string]string)
	initDefaultValues()
}

func initDefaultValues() {
	configurationMap["logger.file"] = "./emm-moneybox.log"
	configurationMap["db.name"] = "emm_moneybox"
	// format: mongodb / mysql
	configurationMap["db.type"] = "mysql"
	// format: mongodb+srv://
	configurationMap["db.mongodb.url"] = os.Getenv("MONGO_DB_URI")
	// format: user:password@tcp(host:port)
	configurationMap["db.mysql.url"] = os.Getenv("MYSQL_DB_URI")
}

func GetConfigByKey(configKey string) string {
	configValue, isExist := configurationMap[configKey]
	if isExist {
		return configValue
	} else {
		return ""
	}
}

func SetConfigByKey(configKey, configValue string) {
	configurationMap[configKey] = configValue
}
