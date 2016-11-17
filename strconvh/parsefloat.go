package strconvh

import "strconv"

// ParseFloat32 interprets a string s in 10-base and returns the corresponding value f (float32) and error.
func ParseFloat32(s string) (float32, error) {
	if valueFloat64, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(valueFloat64), nil
	} else {
		return 0, err
	}
}

// ParseFloat64 interprets a string s in 10-base and returns the corresponding value f (float64) and error.
func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
