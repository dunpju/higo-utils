package timeutil

import (
	"github.com/dunpju/higo-utils/utils/stringutil"
	"regexp"
	"strings"
	"time"
)

var rege = regexp.MustCompile(pattern)

type Timeutil struct {
}

func (this *Timeutil) Time() int64 {
	return Time()
}

func (this *Timeutil) GmtTime(gmt string) int64 {
	return GmtTime(gmt)
}

// 时间戳转时间字符串,格式：Y-m-d H:i:s
func (this *Timeutil) Date(ts int64, format ...string) string {
	return Date(ts, format...)
}

// 纳秒
func (this *Timeutil) Nanoseconds(LowDateTime uint32, HighDateTime uint32) int64 {
	return Nanoseconds(LowDateTime, HighDateTime)
}

//"-2day +1 hour +1 minute"
func (this *Timeutil) Strtotime(datetime string, baseTimestamp ...int64) int64 {
	return Strtotime(datetime, baseTimestamp...)
}

// 当前时间戳
func Time() int64 {
	return time.Now().Unix()
}

const (
	pattern    = `(\+|\-|\s*)(\d+)\s*(year|mouth|day|hour|minute|second)`
	timeLayout = "2006-01-02 15:04:05"
	y          = "Y"
	m          = "m"
	d          = "d"
	h          = "H"
	i          = "i"
	s          = "s"
	minus      = "-"
	second     = "second"
	minute     = "minute"
	hour       = "hour"
	day        = "day"
	mouth      = "mouth"
	year       = "year"
)

func GmtTime(gmt string) int64 {
	if ok, err := regexp.Match(`\d+T\d+:`, []byte(gmt)); !ok || err != nil {
		if !ok {
			panic("format error")
		}
		if err != nil {
			panic(err)
		}
	}
	if ok, err := regexp.Match(`\+08:00$`, []byte(gmt)); !ok || err != nil {
		if !ok {
			panic("format error")
		}
		if err != nil {
			panic(err)
		}
	}
	gmt = strings.Replace(gmt, "T", " ", 1)
	gmt = strings.Replace(gmt, "+08:00", "", 1)
	return Strtotime(gmt)
}

// 时间戳转时间字符串,格式：Y-m-d H:i:s
func Date(ts int64, format ...string) string {
	f := timeLayout
	if len(format) > 0 {
		f = ""
		for _, chr := range format[0] {
			if string(chr) == y {
				f += "2006"
			} else if string(chr) == m {
				f += "01"
			} else if string(chr) == d {
				f += "02"
			} else if string(chr) == h {
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

//"-2day +1 hour +1 minute"
func Strtotime(datetime string, baseTimestamp ...int64) int64 {
	ts := time.Now().Unix()
	if len(baseTimestamp) > 0 {
		ts = baseTimestamp[0]
	}
	if datetime != "" {
		if matchs := rege.FindAllStringSubmatch(datetime, -1); len(matchs) > 0 {
			var add time.Duration
			for _, match := range matchs {
				if minute == match[3] {
					add = time.Minute * time.Duration(stringutil.Int64(match[2]))
					if minus == match[1] {
						add = -add
					}
				} else if hour == match[3] {
					add = time.Hour * time.Duration(stringutil.Int64(match[2]))
					if minus == match[1] {
						add = -add
					}
				} else if day == match[3] {
					add = time.Hour * time.Duration(24*stringutil.Int64(match[2]))
					if minus == match[1] {
						add = -add
					}
				} else {
					add = time.Second * time.Duration(stringutil.Int64(match[2]))
					if minus == match[1] {
						add = -add
					}
				}
				ts = time.Unix(ts, 0).Add(add).Unix()
			}
		} else {
			loc, err := time.LoadLocation("Local")
			if err != nil {
				panic(err)
			}
			theTime, err := time.ParseInLocation(timeLayout, datetime, loc)
			if err != nil {
				panic(err)
			}
			ts = theTime.Unix()
		}
	}
	return ts
}
