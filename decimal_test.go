package gomini

import "testing"

func TestDecimalAddToFloat(t *testing.T) {
	var float64Slice = [][]float64{
		[]float64{2.33,3.99,6.32},
		[]float64{3.33,4.99,8.32},
		[]float64{10.99,4.981,15.971},
		[]float64{2,3,                     5},
		[]float64{2454495034,3451204593,   5905699627},
		[]float64{24544.950,0.345, 24545.295},
		[]float64{0.1,0.1,                   0.2},
		[]float64{0.1,-0.1,                  0},
		[]float64{0,1.001,                 1.001},
	}
	for _,v:=range float64Slice{
		 f := DecimalAddToFloat(v[0],v[1],3)
		if f != v[2] {
			t.Log(f,v[2])
			t.FailNow()
		}
	}
}
func TestDecimalSubToFloat(t *testing.T) {
	var float64Slice = [][]float64{
		[]float64{2,3,                     -1},
		[]float64{12,3,                    9},
		[]float64{-2,9,                    -11},
		[]float64{2454495034,3451204593,   -996709559},
		[]float64{24544.950,0.345, 24544.605},
		[]float64{.1,-0.1,                  0.2},
		[]float64{.1,0.1,                   0},
		[]float64{0,1.001,                 -1.001},
		[]float64{1.001,0,                 1.001},
		[]float64{2.3,0.3,                  2},
	}
	for _,v:=range float64Slice{
		 f := DecimalSubToFloat(v[0],v[1],3)
		if f != v[2] {
			t.Log(f,v[2])
			t.FailNow()
		}
	}
}

