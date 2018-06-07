package mathh

import (
	"math"
	"testing"
)

/*
Not all integer values may be represented as given float type with adequate precision.
min and max constants in function is a minimum and maximum values for tests.
This values are F(floatT, integerT) and are replaced be replacer to appropriate values.
*/

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old float32	Float32	Nextafter32	"= MinInt16"	"= MaxInt16"	int64	Int64	-1}
//replacer:new float32	Float32	Nextafter32	"= MinInt16"	"= MaxInt16"	int		Int		-1}
//replacer:new float32	Float32	Nextafter32	"= MinInt8"		"= MaxInt8"		int8	Int8	-1}
//replacer:new float32	Float32	Nextafter32	"= MinInt16"	"= MaxInt16"	int16	Int16	-1}
//replacer:new float32	Float32	Nextafter32	"= MinInt16"	"= MaxInt16"	int32	Int32	-1}
//replacer:new float64	Float64	Nextafter	"= MinInt32"	"= MaxInt32"	int64	Int64	-1}
//replacer:new float64	Float64	Nextafter	"= MinInt32"	"= MaxInt32"	int		Int		-1}
//replacer:new float64	Float64	Nextafter	"= MinInt8" 	"= MaxInt8" 	int8	Int8	-1}
//replacer:new float64	Float64	Nextafter	"= MinInt16"	"= MaxInt16"	int16	Int16	-1}
//replacer:new float64	Float64	Nextafter	"= MinInt32"	"= MaxInt32"	int32	Int32	-1}
//replacer:new float32	Float32	Nextafter32	"= MinUint16"	"= MaxUint16"	uint64	Uint64	0}
//replacer:new float32	Float32	Nextafter32	"= MinUint16"	"= MaxUint16"	uint	Uint	0}
//replacer:new float32	Float32	Nextafter32	"= MinUint8"	"= MaxUint8"	uint8	Uint8	0}
//replacer:new float32	Float32	Nextafter32	"= MinUint16"	"= MaxUint16"	uint16	Uint16	0}
//replacer:new float32	Float32	Nextafter32	"= MinUint16"	"= MaxUint16"	uint32	Uint32	0}
//replacer:new float64	Float64	Nextafter	"= MinUint32"	"= MaxUint32"	uint64	Uint64	0}
//replacer:new float64	Float64	Nextafter	"= MinUint32"	"= MaxUint32"	uint	Uint	0}
//replacer:new float64	Float64	Nextafter	"= MinUint8"	"= MaxUint8"	uint8	Uint8	0}
//replacer:new float64	Float64	Nextafter	"= MinUint16"	"= MaxUint16"	uint16	Uint16	0}
//replacer:new float64	Float64	Nextafter	"= MinUint32"	"= MaxUint32"	uint32	Uint32	0}

func TestRoundFloat32ToInt64(t *testing.T) {
	const (
		min = MinInt16
		max = MaxInt16
	)

	test := []struct {
		f float32
		i int64
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, -1},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, -1},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter32(min, NegativeInfFloat32()), min},
		{min, min},
		{math.Nextafter32(min, PositiveInfFloat32()), min},
		{math.Nextafter32(max, NegativeInfFloat32()), max},
		{max, max},
		{math.Nextafter32(max, PositiveInfFloat32()), max},

		{NegativeInfFloat32(), MinInt64},
		{PositiveInfFloat32(), MaxInt64},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToInt64(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}
