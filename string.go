package go_string

import (
	"strconv"
	"strings"
)

// ToInt64 字符串转int64
func ToInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

// ToInt 字符串转int
func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}

// ToInt32 字符串转int32
func ToInt32(s string) int32 {
	result, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(result)
}

// ToFloat64 字符串转float64
func ToFloat64(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return result
}

// Join 用分隔符拼接多个字符串
func Join(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

// TrimLeft 从左边删除n个字符 hello 2 --> llo
func TrimLeft(s string, n int) string {
	runes := []rune(s)
	if n >= len(runes) {
		return ""
	}
	return string(runes[n:])
}

// TrimRight 从右边删除n个字符 hello 2 --> hel
func TrimRight(s string, n int) string {
	runes := []rune(s)
	if n >= len(runes) {
		return ""
	}
	return string(runes[:len(runes)-n])
}

// Left 从左边取n个字符 hello 2 --> he
func Left(s string, n int) string {
	runes := []rune(s)
	if n >= len(runes) {
		return s
	}
	return string(runes[:n])
}

// Right 从右边取n个字符 hello 2 --> lo
func Right(s string, n int) string {
	runes := []rune(s)
	if n >= len(runes) {
		return s
	}
	return string(runes[len(runes)-n:])
}

// Trim 去除首尾空格
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// HasPrefix 判断前缀
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 判断后缀
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Contains 判断是否包含
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// IsEmpty 判断是否为空或只有空格
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Replace 替换所有匹配
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Split 按分隔符分割，返回数组
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}
