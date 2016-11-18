package goutilh

import (
	"github.com/apaxa-go/helper/bytesh"
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
			t.Errorf("%v - %v (%v). Got error: %v", fName, i, kind, err)
		}
		if !v.err && (err == nil) {
			var cmpWith string
			if array {
				cmpWith = v.sArray
			} else {
				cmpWith = v.sSlice
			}
			if string(buf.Bytes()) != cmpWith {
				t.Errorf("%v - %v (%v). Expected: %s, got %s", fName, i, kind, cmpWith, string(buf.Bytes()))
			}
		}
	}
}

func TestWriteBytes(t *testing.T) {
	testWriteBytes(t, WriteBytes, true)
	testWriteBytes(t, WriteBytes, false)
}

func TestWriteBytesStr(t *testing.T) {
	testWriteBytes(t, WriteBytesStr, true)
	testWriteBytes(t, WriteBytesStr, false)
}

func TestWriteBytesAst(t *testing.T) {
	testWriteBytes(t, WriteBytesAst, true)
	testWriteBytes(t, WriteBytesAst, false)
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
