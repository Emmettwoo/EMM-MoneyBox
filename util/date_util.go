package util

import (
	"log"
	"strconv"
	"time"
)

var DEFAULT_DATE_FORMAT_IN_STRING string

func init() {
	DEFAULT_DATE_FORMAT_IN_STRING = "20060102"
}

func FormatDateFromString(dateString string) time.Time {
	date, err := time.Parse(DEFAULT_DATE_FORMAT_IN_STRING, dateString)
	if err != nil {
		log.Fatal(err)
	}
	return date
}

func FormatDateToString(date time.Time) string {
	return date.Format(DEFAULT_DATE_FORMAT_IN_STRING)
}

func FormatDateToStringWithSlash(year, month, day int) string {
	return strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day)
}
