// Package iotesth is a helper for iotest system package.
package iotesth

import "io"

// ErrorWriter returns a Writer that writes to w but stop working after n bytes with error err.
// w may be nil. In that case all underlying writes interprets as good - /dev/null mode.
// If err is nil than io.ErrShortWrite will be used (default value).
func ErrorWriter(w io.Writer, n int64, err error) io.Writer {
	if err == nil {
		err = io.ErrShortWrite
	}
	return &errorWriter{w, n, err}
}

type errorWriter struct {
	w   io.Writer
	n   int64
	err error
}

func (ew *errorWriter) Write(p []byte) (n int, err error) {
	if ew.n <= 0 {
		return 0, ew.err
	}
	// real write
	n = len(p)
	if int64(n) > ew.n {
		n = int(ew.n)
	}

	if ew.w != nil {
		n, err = ew.w.Write(p[:n])
	}
	ew.n -= int64(n)
	//
	if err == nil && n < len(p) {
		err = ew.err
	}
	return
}
