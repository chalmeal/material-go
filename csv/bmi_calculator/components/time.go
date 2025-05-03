package components

import (
	"bmi_calclator/config"
	"fmt"
	"time"
)

// TimeParse
// YYYYMMDD形式の文字列をtime.Time型に変換する
func timeParse(date string) time.Time {
	parseDate, err := time.Parse(config.YYYYMMDD, date)
	if err != nil {
		fmt.Printf("Error parsing date: %v", err)
		return time.Time{}
	}

	return parseDate
}
