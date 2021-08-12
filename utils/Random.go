package utils

import (
	"math/rand"
	"time"
)

type Random struct{}

func init() {
	rand.Seed(time.Now().Unix())
}

func NewRandom() *Random {
	return &Random{}
}

// 随机0-86400
func (this *Random) IntHour24ToSecond() int64 {
	return rand.Int63n(24 * 60 * 60) // 24小时换算成秒
}

// 随机字符串
func (this *Random) String(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机字数字
// max 最大数：eg 1000
func (this *Random) Int(max int) int {
	return rand.Intn(max) + 1
}

func (this *Random) BetweenInt(min int, max int) int {
	if min > max {
		panic("min should not gte max")
	}
again:
	n := rand.Intn(max) + 1
	if n < min {
		goto again
	}
	return n
}

func (this *Random) Int64(max int64) int64 {
	return rand.Int63n(max) + 1
}

func (this *Random) BetweenInt64(min int64, max int64) int64 {
	if min > max {
		panic("min should not gte max")
	}
again:
	n := rand.Int63n(max) + 1
	if n < min {
		goto again
	}
	return n
}
