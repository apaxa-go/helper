package ioh

import (
	"github.com/apaxa-go/helper/bytesh"
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
	testElement{
		[]byte{0x01, 0x02, 0x03},
		"[]byte{0x1, 0x2, 0x3}",
		"[3]byte{0x1, 0x2, 0x3}",
		false,
	},

	// 1
	testElement{
		[]byte{},
		"[]byte{}",
		"[0]byte{}",
		false,
	},

	// 2
	testElement{
		[]byte{10, 20, 30},
		"[]byte{0xa, 0x14, 0x1e}",
		"[3]byte{0xa, 0x14, 0x1e}",
		false,
	},

	// 3
	testElement{
		[]byte{0, 0, 0, 0},
		"[]byte{0x0, 0x0, 0x0, 0x0}",
		"[4]byte{0x0, 0x0, 0x0, 0x0}",
		false,
	},
}

func TestBytesToGoSliceSource(t *testing.T) {
	for i, v := range test {
		buf := bytesh.NewBufferDetail(0, 1)
		err := BytesToGoSliceSource(v.b, buf)
		if (err != nil) != v.err {
			t.Errorf("TestBytesToGoSliceSource - %v. Got error: %v", i, err)
		}
		if !v.err && (err == nil) {
			if v.sSlice != string(buf.Bytes()) {
				t.Errorf("TestBytesToGoSliceSource - %v. Expected: %s, got %s", i, v.sSlice, string(buf.Bytes()))
			}
		}
	}
}

func TestBytesToGoArraySource(t *testing.T) {
	for i, v := range test {
		buf := bytesh.NewBufferDetail(0, 128)
		err := BytesToGoArraySource(v.b, buf)
		if (err != nil) != v.err {
			t.Errorf("TestBytesToGoArraySource - %v. Got error: %v", i, err)
		}
		if !v.err && (err == nil) {
			if v.sArray != string(buf.Bytes()) {
				t.Errorf("TestBytesToGoArraySource - %v. Expected: %s, got %s", i, v.sArray, string(buf.Bytes()))
			}
		}
	}
}
