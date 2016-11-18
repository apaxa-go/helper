// Package binaryh provides some helper functions and types to simplify working with binary package.
package binaryh

import (
	"encoding/binary"
	"io"
)

// Read is binary.Read but with variable number of arguments.
// Read stops reading after first error.
// Returns number of successfully read arguments and error.
func Read(r io.Reader, order binary.ByteOrder, data ...interface{}) (int, error) {
	for i, d := range data {
		if err := binary.Read(r, order, d); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

// Write is binary.Write with variable number of arguments.
// Read stops writing after first error.
// Returns number of successfully wrote arguments and error.
func Write(w io.Writer, order binary.ByteOrder, data ...interface{}) (int, error) {
	for i, d := range data {
		if err := binary.Write(w, order, d); err != nil {
			return i, err
		}
	}
	return len(data), nil
}
