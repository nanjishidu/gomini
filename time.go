package gomini

import (
	"strings"
	"time"

	"github.com/jinzhu/now"
)

var WeekdayString = []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}

// GetWeekdayCnName
// 获取周一到周日
func GetWeekdayCnName(value string) (string, error) {
	tpDate, err := now.ParseInLocation(time.Local, value)
	if err != nil {
		return "", err
	}
	return WeekdayString[tpDate.Weekday()], nil
}

// WebTime
func WebTime(t time.Time) string {
	ftime := t.Format(time.RFC1123)
	if strings.HasSuffix(ftime, "UTC") {
		ftime = ftime[0:len(ftime)-3] + "GMT"
	}
	return ftime
}

// GetTimeAgo
func GetTimeAgo(t int64) (s string) {
	tt := time.Now().Unix() - t

	if tt < 60 {
		s = GetInt64Str(tt) + "秒以前"
	} else if tt < 3600 {
		m := tt / 60
		s = GetInt64Str(m) + "分钟以前"
	} else if tt < 86400 {
		m := tt / 3600
		s = GetInt64Str(m) + "小时以前"
	} else if tt < 2592000 {
		m := tt / 86400
		s = GetInt64Str(m) + "天以前"
	} else if tt < 2592000*12 {
		m := tt / 2592000
		s = GetInt64Str(m) + "月以前"
	} else {
		m := tt / (2592000 * 12)
		s = GetInt64Str(m) + "年以前"
	}
	return
}

const (
	DateTypeYear = iota
	DateTypeMonth
	DateTypeWeekday
	DateTypeDay
)

// GetDateBetweenByDateTypeAndDateTimeAndTimeRange
func GetDateBetweenByDateTypeAndDateTimeAndTimeRange(dateType int, dateTime string, timeFormatParams ...string) (string, string, [][]string) {
	startTime, endTime := getDateBetweenByDateTypeAndDateTime(dateType, dateTime)

	var timeFormat = "2006-01-02"
	if len(timeFormatParams) > 0 {
		timeFormat = timeFormatParams[0]
	}
	var (
		startTimeStr   = startTime.Format(timeFormat)
		endTimeStr     = endTime.Format(timeFormat)
		timeRangeSlice [][]string
	)
	for {
		dateTimeSlice := []string{startTime.Format(timeFormat)}
		switch dateType {
		case DateTypeYear:
			startTime = startTime.AddDate(0, 1, 0)
		case DateTypeMonth:
			startTime = startTime.AddDate(0, 0, 1)
		case DateTypeWeekday:
			startTime = startTime.AddDate(0, 0, 1)
		case DateTypeDay:
			startTime = startTime.Add(time.Hour)
		}
		if endTime.Before(startTime) {
			break
		}
		dateTimeSlice = append(dateTimeSlice, startTime.Format(timeFormat))
		timeRangeSlice = append(timeRangeSlice, dateTimeSlice)
	}
	return startTimeStr, endTimeStr, timeRangeSlice
}

// GetDateBetweenByDateTypeAndDateTime
func GetDateBetweenByDateTypeAndDateTime(dateType int, dateTime string, timeFormatParams ...string) (startTimeStr, endTimeStr string) {
	startTime, endTime := getDateBetweenByDateTypeAndDateTime(dateType, dateTime)
	var timeFormat = "2006-01-02"
	if len(timeFormatParams) > 0 {
		timeFormat = timeFormatParams[0]
	}
	return startTime.Format(timeFormat), endTime.Format(timeFormat)
}

// getDateBetweenByDateTypeAndDateTime
func getDateBetweenByDateTypeAndDateTime(dateType int, dateTime string) (startTime, endTime time.Time) {
	t := now.MustParseInLocation(time.Local, dateTime)
	switch dateType {
	case DateTypeYear:
		startTime = now.With(t).BeginningOfYear()
		endTime = now.With(t.AddDate(1, 0, 0)).BeginningOfYear()
	case DateTypeMonth:
		startTime = now.With(t).BeginningOfMonth()
		endTime = now.With(t.AddDate(0, 1, 0)).BeginningOfMonth()
	case DateTypeWeekday:
		now.WeekStartDay = time.Monday
		startTime = now.With(t).BeginningOfWeek()
		endTime = startTime.AddDate(0, 0, 7)
	case DateTypeDay:
		startTime = now.With(t).BeginningOfDay()
		endTime = now.With(t.AddDate(0, 0, 1)).BeginningOfDay()
	}
	return
}
