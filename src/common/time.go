package common

import "time"

var (
	TimeLocation, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
)

func GetVietNamTimeMillisecond() int64 {
	return time.Now().In(TimeLocation).UnixNano() / int64(time.Millisecond)
}

func GetVietNamTime() int64 {
	return time.Now().In(TimeLocation).Unix()
}
