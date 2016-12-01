package iotesth

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestErrorWriter(t *testing.T) {
	data := make([]byte, 1024)

	cn := 0
	ew := ErrorWriter(nil, int64(cn), nil)
	if n, err := ew.Write(data); n != cn || err != io.ErrShortWrite {
		t.Errorf("Expect %v & %v, got %v & %v", cn, io.ErrShortWrite, n, err)
	}

	cerr := errors.New("Custom error")
	cn = 10
	ew = ErrorWriter(nil, int64(cn), cerr)
	if n, err := ew.Write(data); n != cn || err != cerr {
		t.Errorf("Expect %v & %v, got %v & %v", cn, cerr, n, err)
	}

	cn = 100
	ew = ErrorWriter(nil, int64(cn), cerr)
	if n, err := ew.Write(data[:cn/2]); n != cn/2 || err != nil {
		t.Errorf("Expect %v & %v, got %v & %v", cn/2, nil, n, err)
	}
	cn -= cn / 2
	if n, err := ew.Write(data[:cn]); n != cn || err != nil {
		t.Errorf("Expect %v & %v, got %v & %v", cn, nil, n, err)
	}
	if n, err := ew.Write(data[:1]); n != 0 || err != cerr {
		t.Errorf("Expect %v & %v, got %v & %v", 0, cerr, n, err)
	}
	if n, err := ew.Write(data); n != 0 || err != cerr {
		t.Errorf("Expect %v & %v, got %v & %v", 0, cerr, n, err)
	}

	cw := &bytes.Buffer{}
	cn = 100
	ew = ErrorWriter(cw, int64(cn), cerr)
	if n, err := ew.Write(data); n != cn || err != cerr {
		t.Errorf("Expect %v & %v, got %v & %v", cn, cerr, n, err)
	}
	if cwl := cw.Len(); cwl != cn {
		t.Errorf("Expect %v, got %v", cn, cwl)
	}
}
