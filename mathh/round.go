package mathh

// Round returns nearest int64 for given float64
func Round(f float64) int64 {
	switch {
	case f < 0.5 && f > -0.5:
		return 0
	case f > 0:
		return int64(f + 0.5)
	default:
		return int64(f - 0.5)
	}
}
