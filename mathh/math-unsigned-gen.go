//replacer:generated-file
package mathh


// DivideRoundUint divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint(a, b uint) (c uint) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// PowUint returns a**b (a raised to power b).
// Warning: where is no any check for overflow.
func PowUint(a, b uint) uint {
	p := uint(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowModUint computes a**b mod m (modular integer power) using binary powering algorithm.
func PowModUint(a, b, m uint) uint {
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

// DivideRoundUint8 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint8(a, b uint8) (c uint8) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// PowUint8 returns a**b (a raised to power b).
// Warning: where is no any check for overflow.
func PowUint8(a, b uint8) uint8 {
	p := uint8(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowModUint8 computes a**b mod m (modular integer power) using binary powering algorithm.
func PowModUint8(a, b, m uint8) uint8 {
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

// DivideRoundUint16 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint16(a, b uint16) (c uint16) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// PowUint16 returns a**b (a raised to power b).
// Warning: where is no any check for overflow.
func PowUint16(a, b uint16) uint16 {
	p := uint16(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowModUint16 computes a**b mod m (modular integer power) using binary powering algorithm.
func PowModUint16(a, b, m uint16) uint16 {
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

// DivideRoundUint32 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint32(a, b uint32) (c uint32) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// PowUint32 returns a**b (a raised to power b).
// Warning: where is no any check for overflow.
func PowUint32(a, b uint32) uint32 {
	p := uint32(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowModUint32 computes a**b mod m (modular integer power) using binary powering algorithm.
func PowModUint32(a, b, m uint32) uint32 {
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
