package helper
import (
	"strings"
	"strconv"
	"time"
	"fmt"
	"reflect"
)

const layout = "Jan 2, 2006"

func CountPublicationDateByHoursAgo(hoursAgoStr string) string {
	splitHours := strings.Split(hoursAgoStr, " ")
	fmt.Println(splitHours)
	hoursOnlyStr :="+" + strings.Trim(splitHours[0], " ");
	//hoursOnlyStr = "17"
	fmt.Println("type", reflect.TypeOf(hoursOnlyStr))
	fmt.Println("hoursOnlyStr", hoursOnlyStr)

	hoursAgo, err := strconv.Atoi(hoursOnlyStr)
	fmt.Println("converting results", hoursAgo, err)
	hoursAgo = -hoursAgo
	fmt.Println("hoursAgo", hoursAgo)
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
	minutesAgo, _ := strconv.Atoi(strings.Trim(splitMinutes[0], " "))
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
