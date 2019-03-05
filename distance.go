package gomini

import "math"

// 地球半径
const earthRadius = 6378137

type GeoPosition struct {
	// Longitude 经度
	Longitude float64
	// Latitude 纬度
	Latitude float64
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
	s := 2.0 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+
		math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * earthRadius
	s = math.Floor(s*100) / 100
	return int(s)
}
