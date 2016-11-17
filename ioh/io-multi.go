package ioh

import "io"

/*
TODO Useless as Read can return only part of data. May be FullMultiRead?
// Read perform io.Reader.Read but with variable number of arguments.
// Read stops reading after first error.
// Returns number of successfully read arguments and error.
func Read(r io.Reader, data ...[]byte) (int, error) {
	for i, d := range data {
		if n, err := r.Read(d); err != nil {
			return i, err
		}
	}
	return len(data), nil
}
*/

// Write is io.Writer.Write but with variable number of arguments.
// Write stops writing after first error.
// Returns number of successfully wrote arguments and error.
func Write(w io.Writer, data ...[]byte) (int, error) {
	for i, d := range data {
		if _, err := w.Write(d); err != nil {
			return i, err
		}
	}
	return len(data), nil
}
