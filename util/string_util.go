package util

import (
	"strconv"
)

func ToInteger(origin string) int {
	toInteger, err := strconv.Atoi(origin)
	if err != nil {
		Logger.Errorln(err)
	}
	return toInteger
}
