package util

import (
	"time"
)

//func Time() int64 {
//	return time.Now().Unix() - 5*60*60
//}

func Time() int64 {
	return time.Now().Unix()
}

func CurrentDate() string {
	timeLayout := "2006-01-02"
	dateTimeStr := time.Unix(Time(), 0).Format(timeLayout)
	return dateTimeStr
}

func CurrentDateTime() string {
	timeLayout := "2006-01-02 15:04:05"
	dateTimeStr := time.Unix(time.Now().Unix(), 0).Format(timeLayout)
	return dateTimeStr
}

func MilliSecondTime() int64 {
	return int64(time.Now().UnixNano() / 1e6)
}

func NewYorkTime() int64 {
	format := "2006-01-02 15:04:05"
	day := time.Now().Format("2006-01-02 15:04:05")
	local, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation(format, day, local)
	return t.UnixNano() / 1e9
	//t.UnixNano() / 1e6 毫秒
}

func GetDateTime(date string) (int64, int64) {
	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//toDayStartTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	//toDayEndTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Local).Unix()
	//返回当天0点和23点59分的时间戳
	return startTime.Unix(), end.Unix()
}
