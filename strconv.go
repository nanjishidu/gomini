package gomini

import (
	"strconv"
)

func GetStrBool(strv string, def ...bool) (bool, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseBool(strv)
}

func GetStrFloat(strv string, def ...float64) (float64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(strv, 64)
}

func GetStrUint(strv string, def ...uint) (uint, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	a, err := strconv.Atoi(strv)
	return uint(a), err
}

func GetStrInt(strv string, def ...int) (int, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	a, err := strconv.Atoi(strv)
	return int(a), err
}
func GetStrInt8(strv string, def ...int8) (int8, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	a, err := strconv.ParseInt(strv, 10, 8)
	return int8(a), err
}

func GetStrUint8(strv string, def ...uint8) (uint8, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	ret, err := strconv.ParseUint(strv, 10, 8)
	return uint8(ret), err
}
func GetStrInt16(strv string, def ...int16) (int16, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	a, err := strconv.ParseInt(strv, 10, 16)
	return int16(a), err
}
func GetStrUint16(strv string, def ...uint16) (uint16, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	ret, err := strconv.ParseUint(strv, 10, 16)
	return uint16(ret), err
}
func GetStrInt32(strv string, def ...int32) (int32, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	a, err := strconv.ParseInt(strv, 10, 32)
	return int32(a), err
}
func GetStrUint32(strv string, def ...uint32) (uint32, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	ret, err := strconv.ParseUint(strv, 10, 32)
	return uint32(ret), err
}

func GetStrInt64(strv string, def ...int64) (int64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	ret, err := strconv.ParseInt(strv, 10, 64)
	return ret, err
}
func GetStrUint64(strv string, def ...uint64) (uint64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	ret, err := strconv.ParseUint(strv, 10, 64)
	return ret, err
}
func GetStrFloat64(strv string, def ...float64) (float64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	ret, err := strconv.ParseFloat(strv, 64)
	return ret, err
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
