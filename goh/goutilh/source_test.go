package goutilh

import (
	"github.com/apaxa-go/helper/bytesh"
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/testingh/iotesth"
	"io"
	"reflect"
	"runtime"
	"testing"
)

type testElement struct {
	b      []byte
	sSlice string
	sArray string
	err    bool
}

var test = []testElement{
	// 0
	{
		[]byte{0x01, 0x02, 0x03},
		"[]byte{1, 2, 3}",
		"[3]byte{1, 2, 3}",
		false,
	},

	// 1
	{
		[]byte{},
		"[]byte{}",
		"[0]byte{}",
		false,
	},

	// 2
	{
		[]byte{10, 20, 30},
		"[]byte{10, 20, 30}",
		"[3]byte{10, 20, 30}",
		false,
	},

	// 3
	{
		[]byte{0, 0, 0, 0},
		"[]byte{0, 0, 0, 0}",
		"[4]byte{0, 0, 0, 0}",
		false,
	},

	// 4
	{
		[]byte{0, 255},
		"[]byte{0, 255}",
		"[2]byte{0, 255}",
		false,
	},
}

func testWriteBytes(t *testing.T, f func([]byte, io.Writer, bool) error, array bool) {
	fName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	var kind string
	if array {
		kind = "array"
	} else {
		kind = "slice"
	}

	for i, v := range test {
		buf := bytesh.NewBufferDetail(0, 1)
		err := f(v.b, buf, array)
		if (err != nil) != v.err {
			t.Errorf("%v - %v (%v): got error: %v", fName, i, kind, err)
		}
		if !v.err && (err == nil) {
			var cmpWith string
			if array {
				cmpWith = v.sArray
			} else {
				cmpWith = v.sSlice
			}
			if string(buf.Bytes()) != cmpWith {
				t.Errorf("%v - %v (%v): expect %s, got %s", fName, i, kind, cmpWith, string(buf.Bytes()))
			}
		}
	}
}

func TestWriteBytes(t *testing.T) {
	testWriteBytes(t, WriteBytes, true)
	testWriteBytes(t, WriteBytes, false)
}

func TestWriteBytes2(t *testing.T) {
	b := make([]byte, astThreshold+1)
	for i := range b {
		b[i] = byte(i % mathh.MaxUint8)
	}

	buf := bytesh.NewBufferDetail(0, 1)
	for _, array := range []bool{false, true} {
		buf.Reset()
		if err := WriteBytes(b, buf, array); err != nil {
			t.Errorf("got error %v", err)
		}
		s := string(buf.Bytes())

		buf.Reset()
		if err := WriteBytesAst(b, buf, array); err != nil {
			t.Errorf("got error %v", err)
		}
		sAst := string(buf.Bytes())

		buf.Reset()
		if err := WriteBytesStr(b, buf, array); err != nil {
			t.Errorf("got error %v", err)
		}
		sStr := string(buf.Bytes())

		if s != sAst || s != sStr {
			t.Errorf("different results: %v %v %v", s, sAst, sStr)
		}
	}
}

func TestWriteBytesStr(t *testing.T) {
	testWriteBytes(t, WriteBytesStr, true)
	testWriteBytes(t, WriteBytesStr, false)
}

func TestWriteBytesAst(t *testing.T) {
	testWriteBytes(t, WriteBytesAst, true)
	testWriteBytes(t, WriteBytesAst, false)
}

func TestWriteBytesStr2(t *testing.T) {
	data := []byte{1, 2, 3}

	for _, i := range []int{1, 7, 9, 14} {
		w := iotesth.ErrorWriter(nil, int64(i), nil)
		if err := WriteBytesStr(data, w, false); err != io.ErrShortWrite {
			t.Errorf("expect %v, got %v", io.ErrShortWrite, err)
		}
	}
}

func benchmarkWriteBytes(b *testing.B, f func([]byte, io.Writer, bool) error) {
	buf := bytesh.NewBufferDetail(0, 1)
	for i := 0; i < b.N; i++ {
		f(test[i%len(test)].b, buf, i%2 == 0)
		buf.Reset()
	}
}

func BenchmarkWriteBytesAst(b *testing.B) {
	benchmarkWriteBytes(b, WriteBytesAst)
}

func BenchmarkWriteBytesStr(b *testing.B) {
	benchmarkWriteBytes(b, WriteBytesStr)
}
