package httpserv

import "strings"

// IncomingHeaderMatcher 控制输入header
func IncomingHeaderMatcher(key string) (string, bool) {
	lowerKey := strings.ToLower(key)
	if lowerKey == "connection" || lowerKey == "keep-alive" {
		return "", false
	}

	// 默认接受所有header
	return key, true

	// 自定义接受header
	// res, ok := runtime.DefaultHeaderMatcher(key)
	// if ok {
	// 	return res, ok
	// }

	// if strings.HasPrefix(key, "xxx-") {
	// 	return key, true
	// }

	// return "", false
}

// OutgoingHeaderMatcher 控制输出header
func OutgoingHeaderMatcher(key string) (string, bool) {
	// 默认输出所有header
	return key, true
}
