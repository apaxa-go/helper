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

// TODO implement other rounding functions (as for signed integer)

// DivideRoundUint8 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint8(a, b uint8) (c uint8) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// TODO implement other rounding functions (as for signed integer)

// DivideRoundUint16 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint16(a, b uint16) (c uint16) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// TODO implement other rounding functions (as for signed integer)

// DivideRoundUint32 divides a to b and round result to nearest.
//   3 / 2 =  2
func DivideRoundUint32(a, b uint32) (c uint32) {
	c = a / b
	if a%b > (b-1)/2 {
		c++
	}
	return
}

// TODO implement other rounding functions (as for signed integer)
