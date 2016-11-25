package mathh

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/helper/tools-replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old uint64	Uint64
//replacer:new uint	Uint
//replacer:new uint8	Uint8
//replacer:new uint16	Uint16
//replacer:new uint32	Uint32
//replacer:new int	Int
//replacer:new int8	Int8
//replacer:new int16	Int16
//replacer:new int32	Int32
//replacer:new int64	Int64

// Min2Uint64 returns minimum of two passed uint64.
func Min2Uint64(a, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}

// Max2Uint64 returns maximum of two passed uint64.
func Max2Uint64(a, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}
