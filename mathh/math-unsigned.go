package mathh

//replacer:ignore
// TODO implement other rounding functions (as for signed integer)
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
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
