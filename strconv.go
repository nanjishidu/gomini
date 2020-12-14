package gomini

import (
	"github.com/spf13/cast"
)

func GetStrBool(strv string, def ...bool) (bool, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToBoolE(strv)
}

func GetStrFloat(strv string, def ...float64) (float64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToFloat64E(strv)
}

func GetStrFloat32(strv string, def ...float32) (float32, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToFloat32E(strv)
}

func GetStrFloat64(strv string, def ...float64) (float64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToFloat64E(strv)
}

func GetStrUint(strv string, def ...uint) (uint, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToUintE(strv)
}

func GetStrInt(strv string, def ...int) (int, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToIntE(strv)
}

func GetStrInt8(strv string, def ...int8) (int8, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToInt8E(strv)
}

func GetStrUint8(strv string, def ...uint8) (uint8, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToUint8E(strv)
}

func GetStrInt16(strv string, def ...int16) (int16, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToInt16E(strv)
}

func GetStrUint16(strv string, def ...uint16) (uint16, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToUint16E(strv)
}

func GetStrInt32(strv string, def ...int32) (int32, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToInt32E(strv)
}

func GetStrUint32(strv string, def ...uint32) (uint32, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToUint32E(strv)
}
func GetStrInt64(strv string, def ...int64) (int64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToInt64E(strv)
}

func GetStrUint64(strv string, def ...uint64) (uint64, error) {
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return cast.ToUint64E(strv)
}

func GetUint64Str(d uint64) string {
	return cast.ToString(d)
}

func GetInt64Str(d int64) string {
	return cast.ToString(d)
}

func GetUint32Str(d uint32) string {
	return cast.ToString(d)
}

func GetInt32Str(d int32) string {
	return cast.ToString(d)
}

func GetUint16Str(d uint16) string {
	return cast.ToString(d)
}

func GetInt16Str(d int16) string {
	return cast.ToString(d)
}

func GetUint8Str(d uint8) string {
	return cast.ToString(d)
}

func GetInt8Str(d int8) string {
	return cast.ToString(d)
}

func GetUintStr(d uint) string {
	return cast.ToString(d)
}

func GetIntStr(d int) string {
	return cast.ToString(d)
}

func GetFloatStr(f float64) string {
	return cast.ToString(f)
}

func GetFloat32Str(f float32) string {
	return cast.ToString(f)
}

func GetFloat64Str(f float64) string {
	return cast.ToString(f)
}
