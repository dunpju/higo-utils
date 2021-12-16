package hostutil

import (
	"regexp"
	"strings"
)

// 获取http地址
func HttpAddr(ipPortRegexp string, requestHost string) string {

	// 正则判断是否是ip:port
	if m, _ := regexp.MatchString(ipPortRegexp, requestHost); !m {
		return requestHost
	}

	// 分割
	requestHostSplit := strings.Split(requestHost, ":")

	return requestHostSplit[0]
}
