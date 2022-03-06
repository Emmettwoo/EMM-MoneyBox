package util

import (
	"strconv"
	"time"
)

var DEFAULT_DATE_FORMAT_IN_STRING string

func init() {
	DEFAULT_DATE_FORMAT_IN_STRING = "20060102"
}

func FormatDateFromString(dateString string) time.Time {
	date, _ := time.Parse(DEFAULT_DATE_FORMAT_IN_STRING, dateString)
	return date
}

func FormatDateToString(date time.Time) string {
	return strconv.Itoa(date.Year()) + date.Format(DEFAULT_DATE_FORMAT_IN_STRING[4:5]) + strconv.Itoa(date.Day())
}
