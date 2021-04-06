package utils

import (
	"math/rand"
	"time"
)

type Random struct{}

func NewRandom() *Random {
	return &Random{}
}

// 随机0-86400
func (this *Random) IntHour24ToSecond() int64 {
	return rand.Int63n(24 * 60 * 60) // 24小时换算成秒
}

// 随机字符串
func (this *Random) String(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
