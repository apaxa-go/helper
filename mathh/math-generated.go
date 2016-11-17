//replacer:generated-file
package mathh

// Does not work for MinInt (because MinInt * -1 = MinInt)
func AbsInt(i int) int {
	return (-2*NegativeInt(i) + 1) * i

	//if i < 0 {
	//	return -i
	//}
	//return i
}

func AbsFixInt(i int) int {
	return MaxInt*EqualInt(i, MinInt) + AbsInt(i)*NotEqualInt(i, MinInt)

	//if i == MinInt {
	//	return MaxInt
	//}
	//return AbsInt(i)
}

func AntiAbsInt(i int) int {
	//return (-2*PositiveInt(i) + 1) * i

	if i > 0 {
		return -i
	}
	return i
}

// DivideRoundInt divide a to b and round result to nearest.
// -3 / -2 =  2
// -3 /  2 = -2
func DivideRoundInt(a, b int) (c int) {
	c = a / b
	delta := AntiAbsInt(a % b)
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

	//return a/b + LessInt(AntiAbsInt(a%b), (AntiAbsInt(b)+1)/2)*(SameSignInt(a, b)*2-1)

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

func DivideRoundFixInt(a, b int) int {
	if a == MinInt && b == -1 {
		return MaxInt
	}
	return DivideRoundInt(a, b)
}

// DivideCeilInt divide a to b and round result to nearest not less number.
// A.k.a. round up, round towards plus infinity.
// -3 / -2 =  2
// -3 /  2 = -1
func DivideCeilInt(a, b int) int {
	return a/b + NotZeroInt(a%b)*SameSignInt(a, b)
}

// a**b, b>=0
func PowInt(a, b int) int {
	p := int(1)
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
func PowModInt(a, b, m int) int {
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

// Does not work for MinInt8 (because MinInt8 * -1 = MinInt8)
func AbsInt8(i int8) int8 {
	return (-2*NegativeInt8(i) + 1) * i

	//if i < 0 {
	//	return -i
	//}
	//return i
}

func AbsFixInt8(i int8) int8 {
	return MaxInt8*EqualInt8(i, MinInt8) + AbsInt8(i)*NotEqualInt8(i, MinInt8)

	//if i == MinInt8 {
	//	return MaxInt8
	//}
	//return AbsInt8(i)
}

func AntiAbsInt8(i int8) int8 {
	//return (-2*PositiveInt8(i) + 1) * i

	if i > 0 {
		return -i
	}
	return i
}

// DivideRoundInt8 divide a to b and round result to nearest.
// -3 / -2 =  2
// -3 /  2 = -2
func DivideRoundInt8(a, b int8) (c int8) {
	c = a / b
	delta := AntiAbsInt8(a % b)
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

	//return a/b + LessInt8(AntiAbsInt8(a%b), (AntiAbsInt8(b)+1)/2)*(SameSignInt8(a, b)*2-1)

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

func DivideRoundFixInt8(a, b int8) int8 {
	if a == MinInt8 && b == -1 {
		return MaxInt8
	}
	return DivideRoundInt8(a, b)
}

// DivideCeilInt8 divide a to b and round result to nearest not less number.
// A.k.a. round up, round towards plus infinity.
// -3 / -2 =  2
// -3 /  2 = -1
func DivideCeilInt8(a, b int8) int8 {
	return a/b + NotZeroInt8(a%b)*SameSignInt8(a, b)
}

// a**b, b>=0
func PowInt8(a, b int8) int8 {
	p := int8(1)
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
func PowModInt8(a, b, m int8) int8 {
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

// Does not work for MinInt16 (because MinInt16 * -1 = MinInt16)
func AbsInt16(i int16) int16 {
	return (-2*NegativeInt16(i) + 1) * i

	//if i < 0 {
	//	return -i
	//}
	//return i
}

func AbsFixInt16(i int16) int16 {
	return MaxInt16*EqualInt16(i, MinInt16) + AbsInt16(i)*NotEqualInt16(i, MinInt16)

	//if i == MinInt16 {
	//	return MaxInt16
	//}
	//return AbsInt16(i)
}

func AntiAbsInt16(i int16) int16 {
	//return (-2*PositiveInt16(i) + 1) * i

	if i > 0 {
		return -i
	}
	return i
}

// DivideRoundInt16 divide a to b and round result to nearest.
// -3 / -2 =  2
// -3 /  2 = -2
func DivideRoundInt16(a, b int16) (c int16) {
	c = a / b
	delta := AntiAbsInt16(a % b)
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

	//return a/b + LessInt16(AntiAbsInt16(a%b), (AntiAbsInt16(b)+1)/2)*(SameSignInt16(a, b)*2-1)

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

func DivideRoundFixInt16(a, b int16) int16 {
	if a == MinInt16 && b == -1 {
		return MaxInt16
	}
	return DivideRoundInt16(a, b)
}

// DivideCeilInt16 divide a to b and round result to nearest not less number.
// A.k.a. round up, round towards plus infinity.
// -3 / -2 =  2
// -3 /  2 = -1
func DivideCeilInt16(a, b int16) int16 {
	return a/b + NotZeroInt16(a%b)*SameSignInt16(a, b)
}

// a**b, b>=0
func PowInt16(a, b int16) int16 {
	p := int16(1)
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
func PowModInt16(a, b, m int16) int16 {
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

// Does not work for MinInt32 (because MinInt32 * -1 = MinInt32)
func AbsInt32(i int32) int32 {
	return (-2*NegativeInt32(i) + 1) * i

	//if i < 0 {
	//	return -i
	//}
	//return i
}

func AbsFixInt32(i int32) int32 {
	return MaxInt32*EqualInt32(i, MinInt32) + AbsInt32(i)*NotEqualInt32(i, MinInt32)

	//if i == MinInt32 {
	//	return MaxInt32
	//}
	//return AbsInt32(i)
}

func AntiAbsInt32(i int32) int32 {
	//return (-2*PositiveInt32(i) + 1) * i

	if i > 0 {
		return -i
	}
	return i
}

// DivideRoundInt32 divide a to b and round result to nearest.
// -3 / -2 =  2
// -3 /  2 = -2
func DivideRoundInt32(a, b int32) (c int32) {
	c = a / b
	delta := AntiAbsInt32(a % b)
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

	//return a/b + LessInt32(AntiAbsInt32(a%b), (AntiAbsInt32(b)+1)/2)*(SameSignInt32(a, b)*2-1)

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

func DivideRoundFixInt32(a, b int32) int32 {
	if a == MinInt32 && b == -1 {
		return MaxInt32
	}
	return DivideRoundInt32(a, b)
}

// DivideCeilInt32 divide a to b and round result to nearest not less number.
// A.k.a. round up, round towards plus infinity.
// -3 / -2 =  2
// -3 /  2 = -1
func DivideCeilInt32(a, b int32) int32 {
	return a/b + NotZeroInt32(a%b)*SameSignInt32(a, b)
}

// a**b, b>=0
func PowInt32(a, b int32) int32 {
	p := int32(1)
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
func PowModInt32(a, b, m int32) int32 {
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
