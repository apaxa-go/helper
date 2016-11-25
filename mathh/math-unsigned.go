package mathh

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32

// DivideRoundUint64 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint64(a, b uint64) (c uint64) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// PowUint64 returns a**b (a raised to power b).
// Warning: where is no any check for overflow.
func PowUint64(a, b uint64) uint64 {
	p := uint64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowModUint64 computes a**b mod m (modular integer power) using binary powering algorithm.
func PowModUint64(a, b, m uint64) uint64 {
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
