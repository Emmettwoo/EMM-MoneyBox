package util

import (
	"reflect"
	"time"
)

var defaultDateFormatInString = "20060102"

func FormatDateFromString(dateString string) time.Time {
	date, err := time.Parse(defaultDateFormatInString, dateString)
	if err != nil {
		Logger.Errorln(err)
	}
	return date
}

func FormatDateToString(date time.Time) string {
	return date.Format(defaultDateFormatInString)
}

func IsDateTimeEmpty(dateTime time.Time) bool {
	return reflect.DeepEqual(dateTime, time.Time{})
}
