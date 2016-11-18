package ioutilh

import "io"

// ReadFull reads up to len(p) bytes into p.
// It is like io.Reader.Read but it tries to read full p.
// It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
// Even if ReadFull returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but less than len(p) bytes, ReadFull waits for more.
// ReadFull finished (on any of):
// - after len(p) bytes will be read (returns <len(p), nil> ),
// - if EOF will be reached (returns <n, nil> ),
// - on any other error (returns <n, err>).
// ReadFull does not return any error if len(p) bytes have been read.
// Warning: ReadFull can waste CPU because of io.Reader.Read may returns zero bytes even without error (depends on Reader implementation).
// In that case ReadFull will not block internally but will try to get new bytes in loop.
func ReadFull(r io.Reader, p []byte) (n int, err error) {
	for err == nil && len(p) > 0 {
		var n2 int
		n2, err = r.Read(p)
		p = p[n2:]
		n += n2
	}

	if err == io.EOF || len(p) == 0 {
		err = nil
	}

	return
}

// Read perform ReadFull but with variable number of arguments.
// It stops if unable to fully read some of p.
// Read returns number of successfully read p's and error.
// It always returns nil error if EOF has been reached or if all p's have been read.
// Where is no way to check how many bytes have been read if not all p's have been read successfully.
func Read(r io.Reader, p ...[]byte) (n int, err error) {
	for i, d := range p {
		var n2 int
		if n2, err = ReadFull(r, d); n2 != len(d) {
			return i, err
		}
	}
	return len(p), nil
}

// Write is like io.Writer.Write but with variable number of arguments.
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
