package helper

import (
	"strings"
	"strconv"
	"time"
)

const layout = "Jan 2, 2006"

func CountPublicationDateByHoursAgo(hoursAgoStr string) string {
	splitHours := strings.Split(hoursAgoStr, " ")
	hoursOnlyStr := strings.Trim(splitHours[0], " ");
	hoursOnlyStr = Utf8ToAcsii(hoursOnlyStr)
	hoursAgo, _ := strconv.Atoi(hoursOnlyStr)
	hoursAgo = -hoursAgo
	t := time.Now().UTC()
	est, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}
	estTime := t.In(est)
	shiftedTime := estTime.Add(time.Duration(hoursAgo) * time.Hour)
	return shiftedTime.Format(layout)
}

func CountPublicationDateByMinutesAgo(minutesAgoStr string) string {
	splitMinutes := strings.Split(minutesAgoStr, " ")
	minutesOnlyStr := strings.Trim(splitMinutes[0], " ")
	minutesOnlyStr = Utf8ToAcsii(minutesOnlyStr)
	minutesAgo, _ := strconv.Atoi(minutesOnlyStr)
	minutesAgo = -minutesAgo
	t := time.Now().UTC()
	est, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}
	estTime := t.In(est)
	shiftedTime := estTime.Add(time.Duration(minutesAgo) * time.Minute)
	return shiftedTime.Format(layout)
}

// function converts unicode number string like "1", "10", "15" to Ascii
// it is to avoid weird error like "strconv.ParseInt: parsing "\u200e1": invalid syntax" for unicode string "1"
func Utf8ToAcsii(utf8Str string) string {
	runes := []rune(utf8Str)
	outString := string(runes[1:(len(runes))])
	return outString
}
