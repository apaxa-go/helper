// Package bytesh provides types for the manipulation of byte slices.
// This package is a helper for bytes base package.
package bytesh

import (
	"errors"
	"github.com/apaxa-go/helper/mathh"
	"io"
	"math"
)

type seekBehaviourT bool

const (
	seekRead  seekBehaviourT = false
	seekWrite                = true
)

const defaultGrowFactor = 1

// Buffer is variable-sized buffer of bytes with Read, Write and Seek methods. Underlying type is slice if bytes.
// It is possible to specify amount of bytes (named GrowFactor) on which underlying slice will be grown if needed. This capability helps to avoid unnecessary increasing underlying slice by 1 byte multiple times which may cause performance gap.
// Buffer implement io.ReadSeeker and io.WriteSeeker - behaviour of Seek method can be changed on runtime by SetSeekRead() and SetSeekWrite().
// APIs of Buffer and bytes.Buffer are similar, main exception is that Buffer currently does not have Rune-related methods and Unread* methods.
type Buffer struct {
	buf              []byte // contents are the bytes buf[off : len(buf)]
	readOff          int    // read at &buf[readOff]
	writeOff         int    // write at &buf[writeOff]
	seekBehaviour    seekBehaviourT
	activeGrowFactor int
}

// GrowFactor returns effective GrowFactor.
func (b *Buffer) GrowFactor() int {
	if b.activeGrowFactor == 0 {
		return defaultGrowFactor
	}
	return b.activeGrowFactor
}

// RemainingBytes returns a slice of the contents of the unread portion of the buffer.
// If the caller changes the contents of the returned slice, the contents of the buffer will change provided there are no intervening method calls on the Buffer.
func (b *Buffer) RemainingBytes() []byte {
	if b.readOff > len(b.buf) {
		return nil
	}
	return b.buf[b.readOff:]
}

// RemainingLen returns the number of bytes of the unread portion of the buffer.
// Usually b.RemainingLen() == len(b.RemainingBytes()). But it is possible to set read position (using Seek[Read]) to yet not written position of buffer (readPosition > Buffer.len) and in that case len(b.RemainingBytes()) will be 0 and b.RemainingLen() will be negative.
func (b *Buffer) RemainingLen() int {
	return len(b.buf) - b.readOff
}

// Bytes returns underlying buffer.
// If the caller changes the contents of the returned slice, the contents of the buffer will change provided there are no intervening method calls on the Buffer.
func (b *Buffer) Bytes() []byte { return b.buf }

// Len returns len of underlying buffer;
// b.Len() == len(b.Bytes()).
func (b *Buffer) Len() int { return len(b.buf) }

// Grow grows the buffer's capacity, if necessary, to guarantee space for another n bytes. After Grow(n), at least n bytes can be written to the buffer without another allocation.
// If n is negative, Grow will panic. If the buffer can't grow it will panic with ErrTooLarge.
func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("Buffer.Grow: negative count")
	}

	if b.writeOff+n > cap(b.buf) {
		k := int(math.Ceil(float64(b.writeOff+n-cap(b.buf)) / float64(b.GrowFactor())))
		buf := make([]byte, len(b.buf), cap(b.buf)+k*b.GrowFactor())
		copy(buf, b.buf)
		b.buf = buf
	}
}

// Next returns a slice containing the next n bytes from the buffer, advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The slice is only valid until the next some call resize underlying buffer.
func (b *Buffer) Next(n int) []byte {
	m := b.RemainingLen()
	// if ReadSeek to nonexistent position m < 0
	if m < 0 {
		return nil
	}
	if n > m {
		n = m
	}
	data := b.buf[b.readOff : b.readOff+n]
	b.readOff += n
	return data
}

// Read reads the next len(p) bytes from the buffer or until the buffer is drained. The return value n is the number of bytes read.
// If the buffer has no data to return, err is io.EOF (unless len(p) is zero); otherwise it is nil.
func (b *Buffer) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	if b.readOff >= len(b.buf) {
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.readOff:])
	b.readOff += n
	return
}

// ReadByte reads and returns the next byte from the buffer. If no byte is available, it returns error io.EOF.
func (b *Buffer) ReadByte() (c byte, err error) {
	if b.readOff >= len(b.buf) {
		return 0, io.EOF
	}
	c = b.buf[b.readOff]
	b.readOff++
	return c, nil
}

// ReadFrom reads data from r until EOF and appends it to the buffer, growing the buffer as needed. The return value n is the number of bytes read.
// Any error except io.EOF encountered during the read is also returned.
// If the buffer becomes too large, ReadFrom will panic with ErrTooLarge.
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	for {
		b.Grow(1)
		m, e := r.Read(b.buf[b.writeOff:cap(b.buf)])
		n += int64(m)
		b.writeOff += m
		b.buf = b.buf[:b.writeOff]
		if e == io.EOF {
			break
		}
		if e != nil {
			return n, e
		}
	}
	return n, nil // err is EOF, so return nil explicitly
}

// Write appends the contents of p to the buffer, growing the buffer as needed. The return value n is the length of p; err is always nil.
// If the buffer becomes too large, Write will panic with ErrTooLarge.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Grow(len(p))
	// increase len of buffer if needed
	if b.writeOff+len(p) > len(b.buf) {
		b.buf = b.buf[:b.writeOff+len(p)]
	}
	n = copy(b.buf[b.writeOff:], p)
	b.writeOff += n
	return n, nil
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// The returned error is always nil, but is included to match bufio.Writer's WriteByte.
// If the buffer becomes too large, WriteByte will panic with ErrTooLarge.
func (b *Buffer) WriteByte(c byte) error {
	b.Grow(1)
	// increase len of buffer if needed
	if b.writeOff+1 > len(b.buf) {
		b.buf = b.buf[:b.writeOff+1]
	}
	b.buf[b.writeOff] = c
	b.writeOff++
	return nil
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of bytes written; it always fits into an int, but it is int64 to match the io.WriterTo interface.
// Any error encountered during the write is also returned.
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	if b.readOff >= len(b.buf) {
		return 0, nil
	}
	m, e := w.Write(b.buf[b.readOff:])
	b.readOff += m
	n = int64(m)
	if e != nil {
		return n, e
	}
	return
}

// PosRead return current read offset (in bytes).
func (b *Buffer) PosRead() int { return b.readOff }

// PosWrite return current write offset (in bytes)
func (b *Buffer) PosWrite() int { return b.writeOff }

// Pos return current read or write offset - Pos share his behaviour with Seek.
// This behaviour can be changed by SetSeekRead and SetSeekWrite methods.
// Technically Pos return PosRead() or PosWrite() result.
func (b *Buffer) Pos() int {
	if b.seekBehaviour == seekRead {
		return b.PosRead()
	}
	return b.PosWrite()
}

// SeekRead change current read position. It is possible (and even does not error) to set position "into the future" (to position yet not written and even overlap end of underlying slice).
// offset define offset in bytes (can be negative).
// whence define offset base and can be 0 (offset from the beginning of buffer), 1 (offset from current read position) and 2 (offset from the end of buffer).
// Error will be returned if combination of offset and whence cause to negative resulting read position (i.e. SeekRead(-1,0)). Error also happened if resulting read position > MaxInt (but offset and result is of int64 type).
// Using unknown whence (not 0, 1, or 2) cause to error.
// SeekRead return new absolute read position (if not error happens) and error.
func (b *Buffer) SeekRead(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case 0:
		abs = offset
	case 1:
		abs = int64(b.readOff) + offset
	case 2:
		abs = int64(len(b.buf)) + offset
	default:
		return 0, errors.New("buffer.Buffer.ReadSeek: invalid whence")
	}
	if abs < 0 || abs > int64(mathh.MaxInt) {
		return 0, errors.New("buffer.Buffer.ReadSeek: negative position")
	}
	b.readOff = int(abs)
	return abs, nil
}

// SeekWrite change current write position. It is analog of SeekRead (see its documentation).
// Additionally to SeekRead behaviour SeekWrite grow buffer if resulting write position require it.
func (b *Buffer) SeekWrite(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case 0:
		abs = offset
	case 1:
		abs = int64(b.writeOff) + offset
	case 2:
		abs = int64(len(b.buf)) + offset
	default:
		return 0, errors.New("buffer.Buffer.WriteSeek: invalid whence")
	}
	if abs < 0 || abs > int64(mathh.MaxInt) {
		return 0, errors.New("buffer.Buffer.WriteSeek: negative position")
	}
	b.writeOff = int(abs)

	//Feel buffer as it was written
	b.Grow(0)
	// this block looks buggy, strange and useless.
	//if b.writeOff > len(b.buf) {
	//	b.buf = b.buf[:b.writeOff]
	//}

	return abs, nil
}

// Seek implements the io.Seeker interface. It call SeekRead or SeekWrite depend on seek behaviour.
// Seek behaviour can be changed by SetSeekRead and SetSeekWrite methods.
func (b *Buffer) Seek(offset int64, whence int) (int64, error) {
	if b.seekBehaviour == seekRead {
		return b.SeekRead(offset, whence)
	}
	return b.SeekWrite(offset, whence)
}

// Reset resets the buffer to be empty and set read and write positions to zero, but it retains the underlying storage for use by future writes.
func (b *Buffer) Reset() {
	b.buf = b.buf[0:0]
	b.readOff, b.writeOff = 0, 0
}

// Cut cuts buffer to toBytes bytes.
// Cut to negative number of bytes or to size greater than Len cause error.
func (b *Buffer) Cut(toBytes int) error {
	if toBytes < 0 {
		return errors.New("cut to negative position")
	}
	if toBytes > len(b.buf) {
		return errors.New("cut to unexisted position")
	}
	b.buf = b.buf[:toBytes]
	return nil
}

// SetSeekRead set Seek behaviour. Subsequents call to Seek will seek for read position.
// SetSeekRead also change Pos behaviour - it will return read position.
func (b *Buffer) SetSeekRead() {
	b.seekBehaviour = seekRead
}

// SetSeekWrite set Seek behaviour. Subsequents call to Seek will seek for write position.
// SetSeekWrite also change Pos behaviour - it will return write position.
func (b *Buffer) SetSeekWrite() {
	b.seekBehaviour = seekWrite
}

// SetGrowFactor change active grow factor.
// Grow factor is amount of bytes on which underlying slice of bytes will be enlarge each time it is required.
// Pass n < 1 cause panic.
func (b *Buffer) SetGrowFactor(n int) {
	if n < 1 {
		panic("buffer.Buffer.Grow: non positive grow factor")
	}
	b.activeGrowFactor = n
}

// NewBuffer create new Buffer using buf as initially underlying slice of bytes.
// Grow factor will be set to 1.
// Seek behaviour will be "read".
func NewBuffer(buf []byte) *Buffer {
	return &Buffer{buf: buf, activeGrowFactor: defaultGrowFactor, seekBehaviour: seekRead}
}

// NewBufferDetail create new Buffer initially capacity of initCapacity and initially grow factor of growFactor.
// Seek behaviour will be "read".
// Passing growFactor < 1 cause panic.
// Passing initCapacity < 0 cause panic.
func NewBufferDetail(initCapacity, growFactor int) (b *Buffer) {
	b = &Buffer{buf: make([]byte, 0, initCapacity), seekBehaviour: seekRead}
	b.Grow(growFactor)
	return
}
