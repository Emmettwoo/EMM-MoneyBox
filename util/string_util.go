package util

import (
	"bytes"
	"strconv"
)

func ToInteger(origin string) int {
	toInteger, err := strconv.Atoi(origin)
	if err != nil {
		Logger.Errorln(err)
	}
	return toInteger
}

func BatchSurroundingWithSingleQuotes(originList []string) []string {
	var targetList = make([]string, len(originList))
	for _, origin := range originList {
		targetList = append(targetList, SurroundingWithSingleQuotes(origin))
	}
	return targetList
}

func SurroundingWithSingleQuotes(origin string) string {
	return "'" + origin + "'"
}

func CombiningWithComma(originList []string) string {

	var targetString = bytes.Buffer{}
	for _, origin := range originList {
		targetString.WriteString(origin)
		targetString.WriteString(", ")
	}
	return targetString.String()
}
