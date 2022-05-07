package util

import "os"

var configurationMap map[string]string

func init() {
	configurationMap = make(map[string]string)
	initDefaultValues()
}

func initDefaultValues() {
	configurationMap["db.name"] = "emm-money-box"
	configurationMap["db.type"] = "mongodb"
	configurationMap["db.url"] = os.Getenv("MONGO_DB_URI")
}

func GetConfigByKey(configKey string) string {
	configValue, isExist := configurationMap[configKey]
	if isExist {
		return configValue
	} else {
		return ""
	}
}

func SetConfigByKey(configKey string, configValue string) {
	configurationMap[configKey] = configValue
}
