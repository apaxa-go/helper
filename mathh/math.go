package mathh

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	Int64
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32

// Does not work for MinInt64 (because MinInt64 * -1 = MinInt64)
func AbsInt64(i int64) int64 {
	return (-2*NegativeInt64(i) + 1) * i

	//if i < 0 {
	//	return -i
	//}
	//return i
}

func AbsFixInt64(i int64) int64 {
	return MaxInt64*EqualInt64(i, MinInt64) + AbsInt64(i)*NotEqualInt64(i, MinInt64)

	//if i == MinInt64 {
	//	return MaxInt64
	//}
	//return AbsInt64(i)
}

func AntiAbsInt64(i int64) int64 {
	//return (-2*PositiveInt64(i) + 1) * i

	if i > 0 {
		return -i
	}
	return i
}

// DivideRoundInt64 divide a to b and round result to nearest.
// -3 / -2 =  2
// -3 /  2 = -2
func DivideRoundInt64(a, b int64) (c int64) {
	c = a / b
	delta := AntiAbsInt64(a % b)
	if b < 0 && delta < (b+1)/2 {
		if a > 0 {
			c--
			return
		}
		c++
		return
	}
	if b > 0 && delta < (-b+1)/2 {
		if a < 0 {
			c--
			return
		}
		c++
		return
	}
	return

	//return a/b + LessInt64(AntiAbsInt64(a%b), (AntiAbsInt64(b)+1)/2)*(SameSignInt64(a, b)*2-1)

	//c = a / b
	//delta := a % b
	//if delta > 0 { // To negative because |MinInt|>MaxInt
	//	delta = -delta
	//}
	//if b < 0 && delta < (b+1)/2 {
	//	if a > 0 {
	//		c--
	//		return
	//	}
	//	c++
	//	return
	//}
	//if b > 0 && delta < (-b+1)/2 {
	//	if a < 0 {
	//		c--
	//		return
	//	}
	//	c++
	//	return
	//}
	//return
}

func DivideRoundFixInt64(a, b int64) int64 {
	if a == MinInt64 && b == -1 {
		return MaxInt64
	}
	return DivideRoundInt64(a, b)
}

// DivideCeilInt64 divide a to b and round result to nearest not less number.
// A.k.a. round up, round towards plus infinity.
// -3 / -2 =  2
// -3 /  2 = -1
func DivideCeilInt64(a, b int64) int64 {
	return a/b + NotZeroInt64(a%b)*SameSignInt64(a, b)
}

// a**b, b>=0
func PowInt64(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Modular integer power: compute a**b mod m using binary powering algorithm
func PowModInt64(a, b, m int64) int64 {
	a = a % m
	p := 1 % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return p
}
