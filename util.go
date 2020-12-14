package gomini

import (
	"strings"
)

// 字符串截取
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// Repeat("?",10,",")
// ?,?,?,?,?,?,?,?,?,?
func Repeat(s string, count int, sep ...string) string {
	var (
		sep2   string
		fields []string
	)
	if len(sep) > 0 {
		sep2 = sep[0]
	} else {
		sep2 = ","
	}
	for i := 0; i < count; i++ {
		fields = append(fields, s)
	}
	return strings.Join(fields, sep2)
}

//Excel 转化
func ConvertToLetter(iCol int) string {
	if iCol <= 0 {
		return ""
	}
	if iCol <= 26 {
		return string(rune(iCol - 1 + 'A'))
	}
	iCol--
	return string(rune(iCol/26%26-1+'A')) + string(rune(iCol%26+'A'))
}
