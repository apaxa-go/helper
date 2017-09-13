package bidi

//go:generate go run ./internal/generator/main.go

const classLengthInBits = 5
const bracketPairShiftShift = -3

func GetClass(r rune) Class           { return data[r] & (1<<classLengthInBits - 1) }
func GetBracketPairShift(r rune) rune { return rune(data[r]>>classLengthInBits) + bracketPairShiftShift }
