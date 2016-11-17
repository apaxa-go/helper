package ioh

import (
	"fmt"
	"io"
	"strconv"
)

// TODO should it be here or in separate package?

func bytesToGoSource(b []byte, prefix string, w io.Writer) error {
	const format = "%#x"

	// Prefix
	_, err := fmt.Fprint(w, prefix)
	if err != nil {
		return err
	}

	// Data
	if len(b) > 0 {
		// First byte
		_, err = fmt.Fprintf(w, format, b[0])
		if err != nil {
			return err
		}

		// All other bytes
		for _, v := range b[1:] {
			_, err = fmt.Fprintf(w, ", "+format, v)
			if err != nil {
				return err
			}
		}
	}

	// Suffix
	_, err = fmt.Fprint(w, "}")

	return err
}

// BytesToGoSliceSource convert given slice of bytes to GoLang source code representation and write it to given io.Writer.
// Example: for b=[]byte{0x01, 0x02, 0x03} BytesToGoSliceSource write "[]byte{0x01, 0x02, 0x03}" AS STRING.
// This function may be useful with "//go:generate" directive to include some (binary) files into app at compile time.
func BytesToGoSliceSource(b []byte, w io.Writer) error {
	return bytesToGoSource(b, "[]byte{", w)
}

// BytesToGoArraySource is similar to BytesToGoSliceSource but generate array instead of slice.
// Example: for b=[]byte{0x01, 0x02, 0x03} BytesToGoSliceSource write "[3]byte{0x01, 0x02, 0x03}" AS STRING.
func BytesToGoArraySource(b []byte, w io.Writer) error {
	return bytesToGoSource(b, "["+strconv.Itoa(len(b))+"]byte{", w)
}
