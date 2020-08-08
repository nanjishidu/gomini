package gomini

import(
	"github.com/shopspring/decimal"
)

//decimal add to string
func DecimalAddToString(f1,f2 float64,round int) string{
	return DecimalAdd(f1,f2).StringFixed(int32(round))
}

//decimal sub to string
func DecimalSubToString(f1,f2 float64,round int)  string{
	return DecimalSub(f1,f2).StringFixed(int32(round))
}

//decimal mul to string
func DecimalMulToString(f1,f2 float64,round int)  string{
	return DecimalMul(f1,f2).StringFixed(int32(round))
}

//decimal div to string
func DecimalDivToString(f1,f2 float64,round int)  string{
	return DecimalDiv(f1,f2).StringFixed(int32(round))
}

//decimal add to float
func DecimalAddToFloat(f1,f2 float64,round int) float64{
	f,_:=DecimalAdd(f1,f2).Round(int32(round)).Float64()
	return f
}

//decimal sub to float
func DecimalSubToFloat(f1,f2 float64,round int)  float64{
	f,_:=DecimalSub(f1,f2).Round(int32(round)).Float64()
	return f
}

//decimal mul to float
func DecimalMulToFloat(f1,f2 float64,round int)  float64{
	f,_:=DecimalMul(f1,f2).Round(int32(round)).Float64()
	return f
}

//decimal div to float
func DecimalDivToFloat(f1,f2 float64,round int)  float64{
	f,_:=DecimalDiv(f1,f2).Round(int32(round)).Float64()
	return f
}

//decimal add
func DecimalAdd(f1,f2 float64) decimal.Decimal{
	return decimal.RequireFromString(GetFloat64Str(f1)).
		Add(decimal.RequireFromString(GetFloat64Str(f2)))
}

//decimal sub
func DecimalSub(f1,f2 float64)  decimal.Decimal{
	return decimal.RequireFromString(GetFloat64Str(f1)).
		Sub(decimal.RequireFromString(GetFloat64Str(f2)))
}

//decimal mul
func DecimalMul(f1,f2 float64)  decimal.Decimal{
	return decimal.RequireFromString(GetFloat64Str(f1)).
		Mul(decimal.RequireFromString(GetFloat64Str(f2)))
}

//decimal div
func DecimalDiv(f1,f2 float64)  decimal.Decimal{
	return decimal.RequireFromString(GetFloat64Str(f1)).
		Div(decimal.RequireFromString(GetFloat64Str(f2)))
}