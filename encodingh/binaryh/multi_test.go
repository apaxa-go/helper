package binaryh

import (
	"bytes"
	"encoding/binary"
	"github.com/apaxa-go/helper/bytesh"
	"reflect"
	"testing"
)

//TODO divide to different function
func TestRead(t *testing.T) {
	b := []byte{0x01, 0x02, 0x03}
	b1 := make([]byte, 1)
	b2 := make([]byte, 2)
	buf := bytes.NewReader(b)

	// check no arguments
	n := 0
	i, err := Read(buf, binary.LittleEndian)
	if err != nil {
		t.Errorf("TestRead. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestRead. Expected numder of arguments: %v, got: %v", n, i)
	}

	//check 2 arguments
	n = 2
	i, err = Read(buf, binary.LittleEndian, b1, b2)
	if err != nil {
		t.Errorf("TestRead. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestRead. Expected numder of arguments: %v, got: %v", n, i)
	} else if !reflect.DeepEqual(b1, b[:1]) && reflect.DeepEqual(b2, b[1:]) {
		t.Errorf("TestRead. Wrong read. Slices are not equal.\nExpected b1: %v\ngot: %v\nExpected b1: %v\ngot: %v", b1, b[:1], b2, b[1:])
	}

	//check 1 argument
	n = 1
	buf = bytes.NewReader(b2)
	i, err = Read(buf, binary.LittleEndian, b1)
	if err != nil {
		t.Errorf("TestRead. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestRead. Expected numder of arguments: %v, got: %v", n, i)
	} else if !reflect.DeepEqual(b1, b2[:1]) {
		t.Errorf("TestRead. Wrong read. Slices are not equal.\nExpected b1: %v\ngot: %v", b1, b2[:1])
	}

	//check 1000 arguments
	n = 1000
	m := make([]interface{}, n)
	for i := range m {
		m[i] = make([]byte, 1)
	}
	b = make([]byte, n)
	for i := range b {
		b[i] = byte(i)
	}
	buf = bytes.NewReader(b)
	i, err = Read(buf, binary.LittleEndian, m...)
	if err != nil {
		t.Errorf("TestRead. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestRead. Expected numder of arguments: %v, got: %v", n, i)
	}
	for i := 0; i < n; i++ {
		mb := m[i].([]byte)
		if mb[0] != b[i] {
			t.Errorf("TestRead. Wrong read. Slices are not equal.\nExpected b[%v]: %v\ngot: %v", i, b[i], mb[0])
		}
	}

	//check error EOF
	b = make([]byte, 0, 0)
	buf = bytes.NewReader(b)
	_, err = Read(buf, binary.LittleEndian, b2)
	if err == nil {
		t.Errorf("TestRead. Expected error EOF but got nil")
	}
}

func TestWrite(t *testing.T) {
	b1 := []byte{0x04}
	b2 := []byte{0x01, 0x02, 0x03}

	buf := bytesh.NewBufferDetail(0, 1)

	// check no argument
	n := 0
	i, err := Write(buf, binary.LittleEndian)
	if err != nil {
		t.Errorf("TestWrite. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestWrite. Expected numder  of arguments: %v, got: %v", n, i)
	}

	// check 1 arguments
	n = 1
	i, err = Write(buf, binary.LittleEndian, b2)
	if err != nil {
		t.Errorf("TestWrite. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestWrite. Expected numder  of arguments: %v, got: %v", n, i)
	} else if !reflect.DeepEqual(buf.Bytes(), b2) {
		t.Errorf("TestWrite. Wrong write. Slices are not equal.\nExpected b1: %v\ngot: %v", b2, buf.Bytes())
	}

	// check 2 arguments
	n = 2
	buf = bytesh.NewBufferDetail(0, 1)
	i, err = Write(buf, binary.LittleEndian, b1, b2)
	if err != nil {
		t.Errorf("TestWrite. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestWrite. Expected numder  of arguments: %v, got: %v", n, i)
	} else if !reflect.DeepEqual(buf.Bytes(), append(b1, b2...)) {
		t.Errorf("TestWrite. Wrong write. Slices are not equal.\nExpected b1: %v\ngot: %v", buf.Bytes(), append(b1, b2...))
	}

	//check 1000 arguments
	n = 1000
	m := make([]interface{}, n)
	var b []byte
	for i := range m {
		m[i] = []byte{0x01}
		b = append(b, 0x01)
	}
	buf = bytesh.NewBufferDetail(n, 1)
	i, err = Write(buf, binary.LittleEndian, m...)
	if err != nil {
		t.Errorf("TestWrite. Got error: %v", err)
	} else if i != n {
		t.Errorf("TestWrite. Expected numder  of arguments: %v, got: %v", n, i)
	} else if !reflect.DeepEqual(buf.Bytes(), b) {
		t.Errorf("TestWrite. Wrong write. Slices are not equal.\nExpected b1: %v\ngot: %v", buf.Bytes(), b)
	}

	//check error EOF
	buf = bytesh.NewBufferDetail(0, 0)
	_, err = Read(buf, binary.LittleEndian, b2)
	if err == nil {
		t.Errorf("TestRead. Expected error EOF but got nil")
	}
}
