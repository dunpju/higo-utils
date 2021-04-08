package utils

import (
	"time"
)

// 当前时间戳
func Time() int64 {
	return time.Now().Unix()
}

const (
	timeLayout = "2006-01-02 15:04:05"
	Y          = "Y"
	m          = "m"
	d          = "d"
	H          = "H"
	i          = "i"
	s          = "s"
	minute     = 60
	hour       = 60 * minute
	day        = 24 * hour
	mouth      = 30 * day
)

// 转时间,格式：Y-m-d H:i:s
func Date(ts int64, format ...string) string {
	f := timeLayout
	if len(format) > 0 {
		f = ""
		for _, chr := range format[0] {
			if string(chr) == Y {
				f += "2006"
			} else if string(chr) == m {
				f += "01"
			} else if string(chr) == d {
				f += "02"
			} else if string(chr) == H {
				f += "15"
			} else if string(chr) == i {
				f += "04"
			} else if string(chr) == s {
				f += "05"
			} else {
				f += string(chr)
			}
		}
	}
	return time.Unix(ts, 0).Format(f)
}

// 纳秒
func Nanoseconds(LowDateTime uint32, HighDateTime uint32) int64 {
	// 100-nanosecond intervals since January 1, 1601
	nsec := int64(HighDateTime)<<32 + int64(LowDateTime)
	// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
	nsec -= 116444736000000000
	// convert into nanoseconds
	nsec *= 100
	return nsec
}

func Strtotime(datetime string, baseTimestamp ...int64) int64 {
	ts := time.Now().Unix()
	if datetime != "" {
		loc, err := time.LoadLocation("Local")
		if err != nil {
			panic(err)
		}
		theTime, _ := time.ParseInLocation(timeLayout, datetime, loc)
		ts = theTime.Unix()
	}
	if len(baseTimestamp) > 0 {
		ts = time.Unix(If(baseTimestamp[0] > 0, baseTimestamp[0], ts).(int64), 0).Add(time.Second * time.Duration(5)).Unix()
	}
	return ts
}
