// Package types 提供了一些类型转换的方法
package types

import "strconv"

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
