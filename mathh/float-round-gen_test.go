//replacer:generated-file

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

func TestRoundFloat32ToInt(t *testing.T) {
	const (
		min = MinInt16
		max = MaxInt16
	)

	test := []struct {
		f float32
		i int
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

		{NegativeInfFloat32(), MinInt},
		{PositiveInfFloat32(), MaxInt},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToInt(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToInt8(t *testing.T) {
	const (
		min = MinInt8
		max = MaxInt8
	)

	test := []struct {
		f float32
		i int8
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

		{NegativeInfFloat32(), MinInt8},
		{PositiveInfFloat32(), MaxInt8},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToInt8(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToInt16(t *testing.T) {
	const (
		min = MinInt16
		max = MaxInt16
	)

	test := []struct {
		f float32
		i int16
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

		{NegativeInfFloat32(), MinInt16},
		{PositiveInfFloat32(), MaxInt16},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToInt16(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToInt32(t *testing.T) {
	const (
		min = MinInt16
		max = MaxInt16
	)

	test := []struct {
		f float32
		i int32
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

		{NegativeInfFloat32(), MinInt32},
		{PositiveInfFloat32(), MaxInt32},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToInt32(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToInt64(t *testing.T) {
	const (
		min = MinInt32
		max = MaxInt32
	)

	test := []struct {
		f float64
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

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinInt64},
		{PositiveInfFloat64(), MaxInt64},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToInt64(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToInt(t *testing.T) {
	const (
		min = MinInt32
		max = MaxInt32
	)

	test := []struct {
		f float64
		i int
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

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinInt},
		{PositiveInfFloat64(), MaxInt},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToInt(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToInt8(t *testing.T) {
	const (
		min = MinInt8
		max = MaxInt8
	)

	test := []struct {
		f float64
		i int8
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

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinInt8},
		{PositiveInfFloat64(), MaxInt8},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToInt8(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToInt16(t *testing.T) {
	const (
		min = MinInt16
		max = MaxInt16
	)

	test := []struct {
		f float64
		i int16
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

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinInt16},
		{PositiveInfFloat64(), MaxInt16},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToInt16(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToInt32(t *testing.T) {
	const (
		min = MinInt32
		max = MaxInt32
	)

	test := []struct {
		f float64
		i int32
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

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinInt32},
		{PositiveInfFloat64(), MaxInt32},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToInt32(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToUint64(t *testing.T) {
	const (
		min = MinUint16
		max = MaxUint16
	)

	test := []struct {
		f float32
		i uint64
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter32(min, NegativeInfFloat32()), min},
		{min, min},
		{math.Nextafter32(min, PositiveInfFloat32()), min},
		{math.Nextafter32(max, NegativeInfFloat32()), max},
		{max, max},
		{math.Nextafter32(max, PositiveInfFloat32()), max},

		{NegativeInfFloat32(), MinUint64},
		{PositiveInfFloat32(), MaxUint64},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToUint64(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToUint(t *testing.T) {
	const (
		min = MinUint16
		max = MaxUint16
	)

	test := []struct {
		f float32
		i uint
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter32(min, NegativeInfFloat32()), min},
		{min, min},
		{math.Nextafter32(min, PositiveInfFloat32()), min},
		{math.Nextafter32(max, NegativeInfFloat32()), max},
		{max, max},
		{math.Nextafter32(max, PositiveInfFloat32()), max},

		{NegativeInfFloat32(), MinUint},
		{PositiveInfFloat32(), MaxUint},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToUint(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToUint8(t *testing.T) {
	const (
		min = MinUint8
		max = MaxUint8
	)

	test := []struct {
		f float32
		i uint8
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter32(min, NegativeInfFloat32()), min},
		{min, min},
		{math.Nextafter32(min, PositiveInfFloat32()), min},
		{math.Nextafter32(max, NegativeInfFloat32()), max},
		{max, max},
		{math.Nextafter32(max, PositiveInfFloat32()), max},

		{NegativeInfFloat32(), MinUint8},
		{PositiveInfFloat32(), MaxUint8},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToUint8(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToUint16(t *testing.T) {
	const (
		min = MinUint16
		max = MaxUint16
	)

	test := []struct {
		f float32
		i uint16
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter32(min, NegativeInfFloat32()), min},
		{min, min},
		{math.Nextafter32(min, PositiveInfFloat32()), min},
		{math.Nextafter32(max, NegativeInfFloat32()), max},
		{max, max},
		{math.Nextafter32(max, PositiveInfFloat32()), max},

		{NegativeInfFloat32(), MinUint16},
		{PositiveInfFloat32(), MaxUint16},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToUint16(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat32ToUint32(t *testing.T) {
	const (
		min = MinUint16
		max = MaxUint16
	)

	test := []struct {
		f float32
		i uint32
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter32(min, NegativeInfFloat32()), min},
		{min, min},
		{math.Nextafter32(min, PositiveInfFloat32()), min},
		{math.Nextafter32(max, NegativeInfFloat32()), max},
		{max, max},
		{math.Nextafter32(max, PositiveInfFloat32()), max},

		{NegativeInfFloat32(), MinUint32},
		{PositiveInfFloat32(), MaxUint32},
		{NegativeZeroFloat32(), 0},
		{NaNFloat32(), 0},
	}
	for _, v := range test {
		i := RoundFloat32ToUint32(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToUint64(t *testing.T) {
	const (
		min = MinUint32
		max = MaxUint32
	)

	test := []struct {
		f float64
		i uint64
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinUint64},
		{PositiveInfFloat64(), MaxUint64},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToUint64(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToUint(t *testing.T) {
	const (
		min = MinUint32
		max = MaxUint32
	)

	test := []struct {
		f float64
		i uint
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinUint},
		{PositiveInfFloat64(), MaxUint},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToUint(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToUint8(t *testing.T) {
	const (
		min = MinUint8
		max = MaxUint8
	)

	test := []struct {
		f float64
		i uint8
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinUint8},
		{PositiveInfFloat64(), MaxUint8},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToUint8(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToUint16(t *testing.T) {
	const (
		min = MinUint16
		max = MaxUint16
	)

	test := []struct {
		f float64
		i uint16
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinUint16},
		{PositiveInfFloat64(), MaxUint16},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToUint16(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}

func TestRoundFloat64ToUint32(t *testing.T) {
	const (
		min = MinUint32
		max = MaxUint32
	)

	test := []struct {
		f float64
		i uint32
	}{
		{0.5, 1},
		{0, 0},
		{-0.5, 0},
		{0.48999999, 0},
		{-0.4899999, 0},
		{0.599999, 1},
		{-0.599999, 0},
		{99.999999999, 100},
		{99.48999999999999999999999, 99},
		{99.0000000001, 99},

		{math.Nextafter(min, NegativeInfFloat64()), min},
		{min, min},
		{math.Nextafter(min, PositiveInfFloat64()), min},
		{math.Nextafter(max, NegativeInfFloat64()), max},
		{max, max},
		{math.Nextafter(max, PositiveInfFloat64()), max},

		{NegativeInfFloat64(), MinUint32},
		{PositiveInfFloat64(), MaxUint32},
		{NegativeZeroFloat64(), 0},
		{NaNFloat64(), 0},
	}
	for _, v := range test {
		i := RoundFloat64ToUint32(v.f)
		if i != v.i {
			t.Errorf("%v expect %v, got: %v", v.f, v.i, i)
		}
	}
}
