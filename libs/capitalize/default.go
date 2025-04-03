package capitalize

import (
	"strings"
	"unicode"
)

// Capitalize 首字母大写转换函数
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(strings.ToLower(s))  // 将字符串转换为 rune 切片，支持 Unicode
	runes[0] = unicode.ToUpper(runes[0]) // 首字母大写
	return string(runes)
}
