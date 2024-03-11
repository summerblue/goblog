// Package logger 日志相关
package logger

import "log"

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
