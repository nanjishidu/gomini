package gomini

import "testing"

func TestDecimalAddToFloat(t *testing.T) {
	var float64Slice = [][]float64{
		{2.33, 3.99, 6.32},
		{3.33, 4.99, 8.32},
		{10.99, 4.981, 15.971},
		{2, 3, 5},
		{2454495034, 3451204593, 5905699627},
		{24544.950, 0.345, 24545.295},
		{0.1, 0.1, 0.2},
		{0.1, -0.1, 0},
		{0, 1.001, 1.001},
	}
	for _, v := range float64Slice {
		f := DecimalAddToFloat(v[0], v[1], 3)
		if f != v[2] {
			t.Fatal(f, v[2])
		}
	}
}
func TestDecimalSubToFloat(t *testing.T) {
	var float64Slice = [][]float64{
		{2, 3, -1},
		{12, 3, 9},
		{-2, 9, -11},
		{2454495034, 3451204593, -996709559},
		{24544.950, 0.345, 24544.605},
		{.1, -0.1, 0.2},
		{.1, 0.1, 0},
		{0, 1.001, -1.001},
		{1.001, 0, 1.001},
		{2.3, 0.3, 2},
	}
	for _, v := range float64Slice {
		f := DecimalSubToFloat(v[0], v[1], 3)
		if f != v[2] {
			t.Fatal(f, v[2])
		}
	}
}
