package gomini

import "math"

const earthRadius = 6378137 // 地球半径

type GeoPosition struct {
	Longitude float64 //经度
	Latitude  float64 //纬度
}

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}

// 单位：米
func GeoPositionToDistance(from, to GeoPosition) int {
	radLat1 := rad(from.Latitude)
	radLat2 := rad(to.Latitude)
	a := radLat1 - radLat2
	b := rad(from.Longitude) - rad(to.Longitude)
	s := 2.0 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2))) * earthRadius
	s = math.Floor(s*100) / 100
	return int(s)
}
