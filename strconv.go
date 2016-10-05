package gomini

import (
	"strconv"
)

func GetStrUint(d string) uint {
	a, _ := strconv.Atoi(d)
	return uint(a)
}

func GetStrInt(d string) int {
	a, _ := strconv.Atoi(d)
	return a
}

func GetStrUint64(d string) uint64 {
	if d == "" {
		return uint64(0)
	}
	ret, err := strconv.ParseUint(d, 10, 64)
	if err != nil {
		return uint64(0)
	}
	return ret
}

func GetStrInt64(d string) int64 {
	if d == "" {
		return int64(0)
	}
	ret, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return int64(0)
	}
	return ret
}

func GetStrFloat64(s string) float64 {
	if s == "" {
		return float64(0)
	}
	ret, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(0)
	}
	return ret

}

func GetStrUint32(d string) uint32 {
	a, _ := strconv.Atoi(d)
	return uint32(a)
}

func GetUint64Str(d uint64) string {
	return strconv.FormatUint(d, 10)
}

func GetInt64Str(d int64) string {
	return strconv.FormatInt(d, 10)
}

func GetUintStr(d uint) string {
	return strconv.Itoa(int(d))
}

func GetIntStr(d int) string {
	return strconv.Itoa(d)
}
