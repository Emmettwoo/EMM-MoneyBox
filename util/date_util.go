package util

import (
	"strconv"
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

func FormatDateToStringWithSlash(year, month, day int) string {
	return strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day)
}
