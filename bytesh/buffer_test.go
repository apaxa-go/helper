package bytesh

import (
	"bytes"
	//	"math"
	"testing"
	"testing/iotest"
)

//TODO add more tests

func checkSlicesEqual(t *testing.T, testname string, b, v []byte) {
	if (b == nil && v != nil) || (b != nil && v == nil) {
		t.Errorf("%s: slices not equal. One slice is nill.", testname)
	}
	if len(b) != len(v) {
		t.Errorf("%s: slices not equal. Wrong len. Len(bytes)=%v. Len(v)=%v", testname, len(b), len(v))
	} else {
		for i := range b {
			if b[i] != v[i] {
				t.Errorf("%s: Slises not equal!", testname)
			}
		}
	}
}

func checkPos(t *testing.T, testname string, b Buffer, i int) {
	if b.Pos() != i {
		t.Errorf("%s: Wrong Pos. Expected %v, got %v", testname, i, b.Pos())
	}
}

func TestBytes(t *testing.T) {
	var b = Buffer{
		buf: []byte{0x01, 0x02, 0x03},
	}
	p := []byte{0x01, 0x02, 0x03}
	e := b.Bytes()
	checkSlicesEqual(t, "TestBytes", e, p)

	var b1 Buffer
	var p1 []byte
	e1 := b1.Bytes()
	checkSlicesEqual(t, "TestBytes 2", e1, p1)

}

func TestLen(t *testing.T) {
	var b = Buffer{
		buf: []byte{0x01, 0x02, 0x03},
	}
	l := b.Len()
	if l != 3 {
		t.Errorf("\n(1) Expected len:\n%v\ngot:\n%v", 3, l)
	}

	var b1 Buffer
	l1 := b1.Len()
	if l1 != 0 {
		t.Errorf("\n(2) Expected len:\n%v\ngot:\n%v", 0, l1)
	}
}

//TODO split func
func TestRemainingBytes(t *testing.T) {
	var b = Buffer{
		buf:     []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		readOff: 0,
	}
	p := []byte{0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	if n, err := b.Read(make([]byte, 5)); err != nil || n != 5 {
		t.Errorf("TestRemainingBytes. Error while reading: %v. n should be 5 but n=%v", err, n)
	}
	c := b.RemainingBytes()
	checkSlicesEqual(t, "TestRemainingBytes", c, p)
	if n, err := b.SeekRead(10, 0); err != nil || n != 10 {
		t.Errorf("TestRemainingBytes. Error while seeking: %v. n should be 10 but n=%v", err, n)
	}
	p1 := []byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	c1 := b.RemainingBytes()
	checkSlicesEqual(t, "TestRemainingBytes 2", c1, p1)
	if n, err := b.SeekRead(21, 0); err != nil || n != 21 {
		t.Errorf("TestRemainingBytes. Error while seeking: %v. n should be 21 but n=%v", err, n)
	}
	c2 := b.RemainingBytes()
	checkSlicesEqual(t, "TestRemainingBytes 3", c2, make([]byte, 0))
	// check nil
	if n, err := b.SeekRead(22, 0); err != nil || n != 22 {
		t.Errorf("TestRemainingBytes. Error while seekding: %v. n should be 22 but n=%v", err, n)
	}
	c3 := b.RemainingBytes()
	if c3 != nil {
		t.Errorf("Expexted nil, got %v", c3)
	}
	var b4 Buffer
	c4 := b4.RemainingBytes()
	if c4 != nil {
		t.Errorf("Expected nil , got %v", c4)
	}
}

func TestRemainingBytesNil2(t *testing.T) {
	var b Buffer
	c := b.RemainingBytes()
	if c != nil {
		t.Errorf("Expected nil, got error %v", c)
	}
}

//TODO split func
func TestRemainingLen(t *testing.T) {
	var b = Buffer{
		buf: []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
	}
	p := make([]byte, 5)
	if n, err := b.Read(p); err != nil || n != 5 {
		t.Errorf("TestRemainingLen. Error while reading: %v. n should be 5 but n=%v", err, n)
	}
	c := b.RemainingLen()
	if c != 4 {
		t.Errorf("Expected RemainingLen: %v, got %v", 4, c)
	}
	if n, err := b.SeekRead(9, 0); err != nil || n != 9 {
		t.Errorf("TestRemainingLen. Error while seeking: %v. n should be 9 but n=%v", err, n)
	}
	c1 := b.RemainingLen()
	if c1 != 0 {
		t.Errorf("Expected RemainingLen: %v, got %v", 0, c1)
	}
	if n, err := b.SeekRead(12, 0); err != nil || n != 12 {
		t.Errorf("TestRemainingLen. Error while seeking: %v. n should be 12 but n=%v", err, n)
	}
	c2 := b.RemainingLen()
	if c2 != -3 {
		t.Errorf("Expected RemainingLen: %v, got %v", -3, c2)
	}

	var b3 = Buffer{
		buf: []byte{0x01},
	}

	if n, err := b3.ReadByte(); err != nil || n != b3.Bytes()[0] {
		t.Errorf("TestRemainingLen. Error while reading: %v. n should be: %v but n: %v", err, b3.Bytes()[0], n)
	}
	c3 := b3.RemainingLen()
	if c3 != 0 {
		t.Errorf("Expected RemainingLen: %v, got %v", 0, c3)
	}

	var b4 Buffer
	c4 := b4.RemainingLen()
	if c4 != 0 {
		t.Errorf("Expected RemainingLen: %v, got %v", 0, c4)
	}
}

func TestGrowFactor(t *testing.T) {
	var b = Buffer{
		[]byte{0x01, 0x02, 0x03},
		1,
		2,
		false,
		3,
	}
	e := b.GrowFactor()
	if e != 3 {
		t.Errorf("\n(1) Expected activeGrowFactor: %v, got: %v", 3, e)
	}

	var b1 Buffer
	e1 := b1.GrowFactor()
	if e1 != defaultGrowFactor {
		t.Errorf("\n(2) Expected activeGrowFactor: %v, got: %v", defaultGrowFactor, e)
	}

	b2 := NewBufferDetail(1, 10)
	e2 := b2.GrowFactor()
	if e2 != defaultGrowFactor {
		t.Errorf("\n(3) Expected activeGrowFactor: %v, got: %v", defaultGrowFactor, e2)
	}

	b3 := NewBuffer(make([]byte, 2))
	e3 := b3.GrowFactor()
	if e3 != defaultGrowFactor {
		t.Errorf("\n(3) Expected activeGrowFactor: %v, got: %v", defaultGrowFactor, e3)
	}

	b3.SetGrowFactor(20)
	e4 := b3.GrowFactor()
	if e4 != 20 {
		t.Errorf("\n(3) Expected activeGrowFactor: %v, got: %v", 20, e4)
	}
}

func TestGrow(t *testing.T) {
	var b Buffer
	x := []byte{1}
	y := []byte{2}
	for _, startLen := range []int{0, 1000, 10000} {
		xBytes := bytes.Repeat(x, startLen)
		for _, growLen := range []int{0, 1000, 10000} {
			if n, err := b.Write(xBytes); err != nil || n != startLen {
				t.Errorf("TestGrow. Error while writing: %v. Expected n = %v, got %v", err, startLen, n)
			}
			b.Grow(growLen)
			yBytes := bytes.Repeat(y, growLen)
			if n, err := b.Write(yBytes); err != nil || n != growLen {
				t.Errorf("TestGrow. Error while writing: %v. Expected n = %v, got %v", err, growLen, n)
			}
			if cap(b.Bytes()) != startLen+growLen {
				t.Errorf("TestGrow: Wrong cap - cap(b) = %v, startLen+growLen = %v", cap(b.Bytes()), startLen+growLen)
			}
			//t.Errorf("startLen=%v, growLen=%v\nbuf:%v, cap: %v",startLen, growLen, b.Bytes(), cap(b.Bytes()))
			b.Reset()
			b.buf = make([]byte, 0, 0)
		}
	}
}

func TestGrow2(t *testing.T) {
	x := []byte{1}
	for _, growLen := range []int{0, 10, 100} {
		var b Buffer
		xBytes := bytes.Repeat(x, growLen)
		b.Grow(growLen)
		_, err := b.Write(xBytes)
		if err != nil {
			t.Errorf("TestGrow2 - Error while writing to buf %v", err)
		}
		if cap(b.Bytes()) != growLen {
			t.Errorf("TestGrow2 - error with grow. Wrong capacity. Expected: %v, got: %v", growLen, cap(b.Bytes()))
		}
		//t.Errorf("growLen=%v\nbuf:%v, cap: %v", growLen, b.Bytes(), cap(b.Bytes()))
	}
}

func TestGrowNegative(t *testing.T) {
	var b Buffer
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic expected but no panic had happend")
		}
	}()
	b.Grow(-1)
}

func TestNext(t *testing.T) {
	b := []byte{0, 1, 2, 3, 4}
	tmp := make([]byte, 5)
	for i := 0; i <= 5; i++ {
		for j := i; j <= 5; j++ {
			for k := 0; k <= 6; k++ {
				// 0 <= i <= j <= 5; 0 <= k <= 6
				// Check that if we start with a buffer of length j at offset i and ask for Next(k), we get the right bytes.
				buf := NewBuffer(b[0:j])
				n, _ := buf.Read(tmp[0:i])
				if n != i {
					t.Fatalf("TestNext - Read %d returned %d", i, n)
				}
				bb := buf.Next(k)
				want := k
				if want > j-i {
					want = j - i
				}
				if len(bb) != want {
					t.Fatalf("TestNext - in %d,%d: len(Next(%d)) == %d", i, j, k, len(bb))
				}
				for l, v := range bb {
					if v != byte(l+i) {
						t.Fatalf("TestNext - in %d,%d: Next(%d)[%d] = %d, want %d", i, j, k, l, v, l+i)
					}
				}
			}
		}
	}
}

func TestNext2(t *testing.T) {
	var b = Buffer{
		buf:           []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		seekBehaviour: false,
	}
	if n, err := b.Seek(5, 0); err != nil || n != 5 {
		t.Errorf("TestNext2. Seek error: %v. Expeted n: 5, got: %v", err, n)
	}
	c := b.Next(6)
	checkSlicesEqual(t, "(1) TestNext2", c, b.buf[5:11])
	checkPos(t, "() TestNext2", b, 11)
	c1 := b.Next(20)
	checkSlicesEqual(t, "(2) TestNext2", c1, b.buf[11:])
	checkPos(t, "() TestNext2", b, 21)
	c2 := b.Next(6)
	checkSlicesEqual(t, "(3) TestNext2", c2, make([]byte, 0))
	checkPos(t, "() TestNext2", b, 21)

	var b4 = Buffer{
		buf:     []byte{0x01},
		readOff: 10,
	}
	c4 := b4.Next(2)
	if c4 != nil {
		t.Errorf("(5) TestNext2. Expected nil, got %v", c4)
	}
}

func TestNextNegative(t *testing.T) {
	var b Buffer
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic expected but no panic had happend")
		}
	}()
	b.Next(-1)
}

func TestNextNegativeReadOff(t *testing.T) {
	var b = Buffer{
		buf:     []byte{0x01},
		readOff: -10,
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic expected but no panic had happend")
		}
	}()
	b.Next(1)
}

func TestNewBuffer(t *testing.T) {
	p := []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	buf := NewBuffer(p)
	checkSlicesEqual(t, "TestNewBuffer", buf.Bytes(), p)
	if buf.activeGrowFactor != defaultGrowFactor {
		t.Errorf("TestNewBuffer: Expected activeGrowFactor: %v, got %v", defaultGrowFactor, buf.activeGrowFactor)
	}

	var p1 []byte
	buf1 := NewBuffer(p1)
	checkSlicesEqual(t, "TestNewBuffer", buf1.Bytes(), p1)

}

func TestNewBufferDetail(t *testing.T) {
	b := NewBufferDetail(1, 10)
	if cap(b.buf) != 10 {
		t.Errorf("TestNewBufferDetail. Expected capacity %v, got %v", 10, cap(b.buf))
	}
	b1 := NewBufferDetail(20, 10)
	if cap(b1.buf) != 20 {
		t.Errorf("TestNewBufferDetail. Expected capacity %v, got %v", 20, cap(b1.buf))
	}
	b2 := NewBufferDetail(0, 0)
	if cap(b2.buf) != 0 {
		t.Errorf("TestNewBufferDetail. Expected capacity %v, got %v", 0, cap(b2.buf))
	}

	b3 := NewBufferDetail(0, 10)
	if cap(b3.buf) != 10 {
		t.Errorf("TestNewBufferDetail. Expected capacity %v, got %v", 10, cap(b3.buf))
	}
}

// TODO remove func name from all errors
func TestBasicOperations(t *testing.T) {
	var buf Buffer
	p := []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	for i := 0; i < 5; i++ {
		n, err := buf.Write(p[0:1])
		if n != 1 {
			t.Errorf("TestBasicOperations. Wrote 1 byte, but n == %d\n", n)
		}
		if err != nil {
			t.Errorf("TestBasicOperations. Error should always be nil, but err == %s", err)
		}
		checkSlicesEqual(t, "TestBasicOperations (1)", buf.Bytes(), p[0:1])
		if err = buf.WriteByte(p[1]); err != nil {
			t.Errorf("WriteByte error: %v", err)
		}
		checkSlicesEqual(t, "TestBasicOperations (2)", buf.Bytes(), p[0:2])
		n, err = buf.Write(p[2:19])
		if n != 17 {
			t.Errorf("TestBasicOperations. Wrote 17 bytes, but n == %d", n)
		}
		checkSlicesEqual(t, "TestBasicOperations (3)", buf.Bytes(), p[0:19])
		c, err := buf.ReadByte()
		if err != nil {
			t.Error("TestBasicOperations. ReadByte unexpected eof")
		}
		if c != p[0] {
			t.Errorf("TestBasicOperations. ReadByte wrong value c=%v", c)
		}
		buf.Reset()
		_, err = buf.ReadByte()
		if err == nil {
			t.Error("TestBasicOperations. ReadByte unexpected not eof")
		}
	}
}

// Was a bug: used to give EOF reading empty slice at EOF.
func TestReadEmptyAtEOF(t *testing.T) {
	var slice []byte
	b := new(Buffer)
	n, err := b.Read(slice)
	if err != nil {
		t.Errorf("read error: %v", err)
	}
	if n != 0 {
		t.Errorf("wrong count; got %d want 0", n)
	}
}

func TestRead(t *testing.T) {
	var b = Buffer{
		[]byte{0x01, 0x02, 0x03},
		1,
		2,
		false,
		3,
	}
	p := []byte{0x05, 0x06}
	n, err := b.Read(p)
	if err != nil {
		t.Errorf("(1)TestRead. Error got: %v", err)
	}
	if n != len(p) {
		t.Errorf("(1)TestRead. Expected n: %v, got: %v", len(p), n)
	}
	checkSlicesEqual(t, "(1)TestRead", b.buf[1:3], p)

	p1 := make([]byte, 20)
	_, err1 := b.Read(p1)
	if err1 == nil {
		t.Errorf("(2)TestRead. Expected nil, got error: %v", err1)
	}

	p2 := make([]byte, 0, 20)
	n2, err2 := b.Read(p2)
	if err2 != nil {
		t.Errorf("(3)TestRead. Error got: %v", err2)
	}
	if n2 != len(p2) {
		t.Errorf("(3)TestRead. Expected n: %v, got: %v", len(p2), n2)
	}

	var p3 []byte
	n3, err3 := b.Read(p3)
	if err3 != nil {
		t.Errorf("(3)TestRead. Error got: %v", err3)
	}
	if n3 != 0 {
		t.Errorf("(3)TestRead. Expected n: %v, got: %v", 0, n3)
	}

	n4, err4 := b.Read(nil)
	if err4 != nil {
		t.Errorf("(4)TestRead. Error got: %v", err4)
	}
	if n4 != 0 {
		t.Errorf("(4)TestRead. Expected n: %v, got: %v", 0, n4)
	}
}

func TestReadFrom(t *testing.T) {
	var b Buffer
	p := []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	buf := bytes.NewReader(p)
	n, err := b.ReadFrom(buf)
	if err != nil {
		t.Errorf("(1)TestReadFrom. Error got: %v", err)
	}
	if n != int64(len(p)) {
		t.Errorf("(1)TestReadFrom. Expected n: %v, got: %v", len(p), n)
	}
	checkSlicesEqual(t, "(1)TestReadFrom", b.Bytes(), p)

	var b1 Buffer
	var p1 []byte
	buf1 := bytes.NewReader(p1)
	n1, err1 := b1.ReadFrom(buf1)
	if err1 != nil {
		t.Errorf("(2)TestReadFrom. Error got: %v", err1)
	}
	if n1 != 0 {
		t.Errorf("(2)TestReadFrom. Expected n: %v, got: %v", 0, n1)
	}

	buf2 := iotest.TimeoutReader(bytes.NewReader(p))
	n2, err2 := b.ReadFrom(buf2)
	if err2 == nil {
		t.Error("Expected error, got no error")
	}
	if n2 == 0 {
		t.Error("Expected non zero byte read")
	}
}

func TestWrite(t *testing.T) {
	var b Buffer
	p := []byte{0xff, 0x01, 0x02}
	n, err := b.Write(p)
	if err != nil {
		t.Errorf("(1)TestWrite. Error got: %v", err)
	}
	if n != 3 {
		t.Errorf("(1)TestWrite. Expected n: %v, got: %v", 3, n)
	}
	checkSlicesEqual(t, "(1)TestWrite", b.Bytes(), p)
	b.SetSeekWrite()
	checkPos(t, "(1)TestWrite", b, 3)

	var b1 Buffer
	var p1 []byte
	n1, err1 := b1.Write(p1)
	if err1 != nil {
		t.Errorf("(2)TestWrite. Error got: %v", err1)
	}
	if n1 != 0 {
		t.Errorf("(2)TestWrite. Expected n: %v, got: %v", 0, n1)
	}
	checkSlicesEqual(t, "(2)TestWrite", b1.Bytes(), p1)
	b.SetSeekWrite()
	checkPos(t, "(2)TestWrite", b1, 0)
}

func TestWriteTo(t *testing.T) {
	// 1
	var b = Buffer{
		buf:      []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
		readOff:  2,
		writeOff: 3,
		//	seekBehaviour: true,
	}
	var p []byte
	buf := bytes.NewBuffer(p)
	n, err := b.WriteTo(buf)
	if err != nil {
		t.Errorf("(1)TestWriteTo. Error got: %v", err)
	}
	if n != 8 {
		t.Errorf("(1)TestWriteTo. Expected n: %v, got: %v", 8, n)
	}
	checkSlicesEqual(t, "(1)TestWriteTo", buf.Bytes(), b.buf[2:])

	// 2
	var b1 = Buffer{
		buf:           []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
		readOff:       11,
		writeOff:      3,
		seekBehaviour: true,
	}
	var p1 []byte
	buf1 := bytes.NewBuffer(p1)
	n1, err1 := b1.WriteTo(buf1)
	if err1 != nil {
		t.Errorf("(2)TestWriteTo. Error got: %v", err1)
	}
	if n1 != 0 {
		t.Errorf("(2)TestWriteTo. Expected n: %v, got: %v", 0, n1)
	}
	if len(buf1.Bytes()) != 0 {
		t.Errorf("(2)TestWriteTo. Expected len(buf.Bytes()): %v, got: %v", 0, len(buf1.Bytes()))
	}

	// 3
	var b2 = Buffer{
		buf: []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
	}
	p2 := []byte{0x01, 0x02, 0x03}
	buf2 := bytes.NewBuffer(p2)
	n2, err2 := b2.WriteTo(buf2)
	if err2 != nil {
		t.Errorf("(3)TestWriteTo. Error got: %v", err2)
	}
	if n2 != 10 {
		t.Errorf("(3)TestWriteTo. Expected n: %v, got: %v", 10, n2)
	}
	checkSlicesEqual(t, "(3)TestWriteTo2", buf2.Bytes(), append(p2, b2.buf...))

	// 4
	//var b3 = Buffer{
	//	buf: []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
	//}
	//p3 := []byte{0x01, 0x02, 0x03}
	//buf3 := iotest.TruncateWriter(bytes.NewBuffer(p3),2)
	//n3, err3 := b3.WriteTo(buf3)
	//if err3 == nil {
	//	t.Error("(4)TestWriteTo. Expect error, no error got")
	//}
	//if n3 != 2 {
	//	t.Errorf("(4)TestWriteTo. Expected n: %v, got: %v", 2, n3)
	//}
}

func TestReadWriteByte(t *testing.T) {
	var b Buffer
	p := []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}

	for i, v := range p {
		err := b.WriteByte(v)
		if err != nil {
			t.Errorf("err should always be nil, but err == %s", err)
		}
		checkSlicesEqual(t, "TestReadWriteByte", b.Bytes(), p[:i+1])
		c, err := b.ReadByte()
		if c != p[i] {
			t.Error("read wrong byte")
		}
		if err != nil {
			t.Errorf("err should always be nil, but err == %s", err)
		}
	}
}

//TODO split func
func TestReadWrite(t *testing.T) {
	var b Buffer
	p := []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
	pCompare := []byte{0x00, 0x00, 0xff, 0x01, 0x02, 0x03, 0x00, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}

	b.SetSeekWrite()
	if n, err := b.Seek(2, 0); err != nil || n != 2 {
		t.Errorf("Seek error: %v. Expected n: 2, got: %v", err, n)
	}
	n, err := b.Write(p[0:4])
	if n != 4 {
		t.Errorf("wrote 4 byte, but n == %d\n", n)
	}
	if err != nil {
		t.Errorf("err should always be nil, but err == %s", err)
	}
	checkSlicesEqual(t, "TestReadWrite (1)", b.Bytes(), pCompare[0:6])
	checkPos(t, "TestReadWrite (1)", b, 6)
	b.SetSeekRead()
	checkPos(t, "TestReadWrite (2)", b, 0)
	if n, err := b.Seek(3, 0); err != nil || n != 3 {
		t.Errorf("Seek error: %v. Expected n: 3, got: %v", err, n)
	}
	p1 := make([]byte, 2)
	n1, err1 := b.Read(p1)
	if n1 != 2 {
		t.Errorf("read _ byte, but n == %d\n", n1)
	}
	if err1 != nil {
		t.Errorf("err should always be nil, but err == %s", err1)
	}
	checkSlicesEqual(t, "TestReadWrite (2)", p1, pCompare[3:5])
	checkPos(t, "TestReadWrite (3)", b, 5)

	b.SetSeekWrite()
	position := int64(b.Pos() + 1)
	if n, err := b.Seek(int64(b.Pos()+1), 0); err != nil || n != position {
		t.Errorf("Seek error: %v. Expected n: %v, got: %v", err, p, n)
	}
	n2, err2 := b.Write(p[4:])
	if n2 != 17 {
		t.Errorf("wrote 4 byte, but n == %d\n", n2)
	}
	if err2 != nil {
		t.Errorf("err should always be nil, but err == %s", err2)
	}

	checkSlicesEqual(t, "TestReadWrite (3)", b.Bytes(), pCompare)
	checkPos(t, "TestReadWrite (4)", b, 24)

	b.SetSeekRead()
	checkPos(t, "TestReadWrite (5)", b, 5)
	p2 := make([]byte, 17)
	n3, err3 := b.Read(p2)
	if n3 != 17 {
		t.Errorf("read _ byte, but n == %d\n", n3)
	}
	if err3 != nil {
		t.Errorf("err should always be nil, but err == %s", err3)
	}
	checkSlicesEqual(t, "TestReadWrite (3)", b.buf[3:22], append(p1, p2...))
	checkPos(t, "TestReadWrite (6)", b, 22)
	//t.Errorf("\nb: %v\np_c: %v\np: %v\np1: %v\np2: %v",b.buf, p_compare, p, p1, p2)
}

func TestSeekWrite(t *testing.T) {
	type testSeek struct {
		p       []byte
		offset  int64
		whence  int
		n       int64
		fOffset int64
		err     bool
	}

	var test = []testSeek{

		// 0
		// positive test
		{
			[]byte{0x01, 0x02, 0x03},
			12,
			0,
			12,
			15,
			false,
		},

		// 1
		// negative test - invalid position
		{
			[]byte{0x01, 0x02, 0x03},
			-1,
			0,
			0,
			3,
			true,
		},

		// 2
		// negative test - invalid whence
		{
			[]byte{0x01, 0x02, 0x03},
			0,
			3,
			0,
			3,
			true,
		},

		// 3
		//
		{
			[]byte{0x01, 0x02, 0x03},
			2,
			2,
			2,
			5,
			false,
		},

		// 4
		// whence = 1
		{
			[]byte{0x01, 0x02, 0x03},
			2,
			1,
			2,
			5,
			false,
		},
	}

	for i, v := range test {
		buf := NewBufferDetail(0, 4)
		buf.SetSeekWrite()
		n, err := buf.Seek(v.offset, v.whence)
		if (err != nil) != v.err {
			t.Errorf("Test seekWrite - %d\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if n != v.n {
				t.Errorf("Test seekWrite - %d\nExpected n: %v, got: %v", i, v.n, n)
			}
			_, err := buf.Write(v.p)
			if err != nil {
				t.Errorf("Test seekWrite - %d\nError%v", i, err)
			}
			//	checkSlicesEqual(t, "SeekWrite", buf.buf[12:], v.p)
			if int64(buf.Pos()) != v.fOffset {
				t.Errorf("Test seekWrite - %d\n\nExpected offset: %v, got %v", i, v.fOffset, buf.Pos())
			}
		}
	}
}

func TestSeekRead(t *testing.T) {

	type testSeek struct {
		b       Buffer
		capP    int
		offset  int64
		whence  int
		n       int64
		fOffset int64
		err     bool
	}

	var test = []testSeek{

		// 0
		// positive test
		{
			Buffer{
				buf: []byte{0x01, 0x02, 0x03},
			},
			1,
			0,
			0,
			0,
			1,
			false,
		},

		// 1
		//
		{
			Buffer{
				buf: []byte{0x01, 0x02, 0x03},
			},
			10,
			0,
			0,
			0,
			3,
			false,
		},

		// 2
		//
		{
			Buffer{
				buf: []byte{0x01, 0x02, 0x03},
			},
			10,
			1,
			0,
			1,
			3,
			false,
		},

		// 3
		//
		{
			Buffer{
				buf:     []byte{0x01, 0x02, 0x03},
				readOff: 1,
			},
			10,
			0,
			1,
			1,
			3,
			false,
		},

		// 4
		//
		{
			Buffer{
				buf:     []byte{0x01, 0x02, 0x03},
				readOff: 1,
			},
			10,
			1,
			1,
			2,
			3,
			false,
		},

		// 5
		//
		{
			Buffer{
				buf:     []byte{0x01, 0x02, 0x03},
				readOff: 1,
			},
			10,
			-2,
			2,
			1,
			3,
			false,
		},

		// 6
		// negative test
		{
			Buffer{
				buf: []byte{0x01, 0x02, 0x03},
			},
			10,
			-2,
			0,
			0,
			0,
			true,
		},

		// 7
		// negative test
		{
			Buffer{
				buf: []byte{0x01, 0x02, 0x03},
			},
			10,
			0,
			4,
			0,
			0,
			true,
		},

		//TODO why it's commented
		// 8
		// 9223372036854775807
		/*		testSeek{
				Buffer{
					buf: make([]byte, 9223372036854775809, 9223372036854775809),

				},
				10,
				2,
				2,
				0,
				0,
				false,
			},*/
	}

	for i, v := range test {
		p := make([]byte, v.capP)
		v.b.SetSeekRead()
		n, err := v.b.Seek(int64(v.offset), v.whence)
		if (err != nil) != v.err {
			t.Errorf("Test seekRead - %d\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if n != v.n {
				t.Errorf("Test seekRead - %d\nExpected n: %v, got: %v", i, v.n, n)
			}
			_, err1 := v.b.Read(p)
			if err1 != nil {
				t.Errorf("Test seekRead - %d\nError while reading: %v", i, err1)
			}
			if int64(v.b.Pos()) != v.fOffset {
				t.Errorf("Test seekRead - %d\nExpected offset: %v, got %v", i, v.fOffset, v.b.Pos())
			}
		}
	}
}

func TestSetGrowFactor(t *testing.T) {
	var b Buffer
	x := []byte{1}
	for _, growLen := range []int{1, 5, 10} {
		b.SetGrowFactor(growLen)
		xBytes := bytes.Repeat(x, growLen+1)
		//		b.Grow(growLen)
		_, err := b.Write(xBytes)
		if err != nil {
			t.Errorf("TestSetGrowFactor. Expected error: nil, got %v", err)
		}
		if cap(b.Bytes()) != growLen*2 {
			t.Errorf("TestSetGrowFactor. Expected capacity: %v, got: %v", cap(b.Bytes()), growLen)
		}
		//t.Errorf("growLen=%v\nbuf:%v, cap: %v", growLen, b.Bytes(), cap(b.Bytes()))

		b.Reset()
		b.buf = make([]byte, 0, 0)
	}
}

func TestSetGrowFactorNegative(t *testing.T) {
	var b Buffer
	defer func() {
		if r1 := recover(); r1 == nil {
			t.Error("Panic expected but no panic had happend")
		}
	}()
	b.SetGrowFactor(-1)
}

func TestPos(t *testing.T) {
	var b = Buffer{
		buf:      []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
		readOff:  2,
		writeOff: 3,
	}
	b.SetSeekRead()
	checkPos(t, "(1)TestPos. Error read position", b, 2)
	b.SetSeekWrite()
	checkPos(t, "(2)TestPos. Error write position", b, 3)

	var b1 = Buffer{
		buf:      []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
		readOff:  10,
		writeOff: 20,
	}
	if b1.PosRead() != 10 {
		t.Errorf("(3)TestPos. Error read position, expected %v, got %v", 10, b1.PosRead())
	}
	if b1.PosWrite() != 20 {
		t.Errorf("(4)TestPos. Error write position, expected %v, got %v", 20, b1.PosWrite())
	}
}

//TODO add check capacity after fix
func TestReset(t *testing.T) {
	var b = Buffer{
		buf:              []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
		readOff:          2,
		writeOff:         3,
		seekBehaviour:    true,
		activeGrowFactor: 2,
	}
	b.Reset()
	if b.Len() != 0 || b.PosRead() != 0 || b.PosWrite() != 0 {
		t.Errorf("TestReset. Buffer doesn't reset - %v", b)
	}

	var b1 Buffer
	b1.Reset()
	if b1.Len() != 0 || b1.PosRead() != 0 || b1.PosWrite() != 0 {
		t.Errorf("(1)TestReset. Buffer doesn't reset - %v", b1)
	}
}

func TestCut(t *testing.T) {
	var b = Buffer{
		buf: []byte{0xff, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
	}
	err := b.Cut(-1)
	if err == nil {
		t.Error("TestCut. Expected error but got nil")
	}

	err1 := b.Cut(8)
	if err1 == nil {
		t.Errorf("(1)TestCut. Expected error but got nil. %v", b)
	}

	p2 := []byte{0xff, 0x01, 0x02}
	err2 := b.Cut(3)
	if err2 != nil {
		t.Errorf("(2)TestCut. Error %v", err2)
	}
	checkSlicesEqual(t, "(3)TestCut", p2, b.Bytes())

	err3 := b.Cut(0)
	if err3 != nil {
		t.Error("(4)TestCut. Expected error but got nil")
	}
	checkSlicesEqual(t, "(5)TestCut", make([]byte, 0), b.Bytes())
}

/*
func TestLargeByteWrite(t *testing.T) {
	var b Buffer
	p1 := []byte{0xff, 0x01, 0x02}
	p := append(make([]byte, math.MaxInt32-len(p1), math.MaxInt32), p1...)
	n, err := b.Write(p)
	if err != nil {
		t.Errorf("(1)TestLargeByteWrites. Error got: %v", err)
	}
	if n != len(p) {
		t.Errorf("(2)TestLargeByteWrites. Expected n: %v, got: %v", len(p), n)
	}
	checkSlicesEqual(t, "(3)TestWrite", b.Bytes(), p)
	b.SetSeekWrite()
	checkPos(t, "(4)TestWrite", b, len(p))
}

func TestLargeByteRead(t *testing.T) {
	var b Buffer
	p1 := []byte{0xff, 0x01, 0x02}
	b.buf = append(make([]byte, math.MaxInt32-len(p1), math.MaxInt32), p1...)

	p := make([]byte, math.MaxInt32)
	n, err := b.Read(p)
	if err != nil {
		t.Errorf("(1)TestLargeByteRead. Error got: %v", err)
	}
	if n != len(p) {
		t.Errorf("(2)TestLargeByteRead. Expected n: %v, got: %v", len(p), n)
	}
	checkSlicesEqual(t, "(3)TestRead", b.Bytes(), p)
	checkPos(t, "(4)TestRead", b, len(p))
}*/

func BenchmarkSNRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf Buffer
		p := []byte{0xff, 0x01, 0x02}
		if _, err := buf.Write(p); err != nil {
			b.Errorf("Write error: %v", err)
		}
	}
}

func BenchmarkSNWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf Buffer
		p := []byte{0xff, 0x01, 0x02}
		if _, err := buf.Read(p); err != nil {
			b.Errorf("Read error: %v", err)
		}
	}
}
