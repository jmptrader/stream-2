/*
Stream is a golang package to help read/write byte streams which could be easily read/write with java.io.DataInput and java.io.DataOutput.

Borrowing heavily from golang packages such as binary and bytes.  Lots of error have been just ignored. BigEndian only.

*/
package stream

import (
	"encoding/binary"
	"bytes"
)


type Reader struct {
	buf []byte
	off int  //current reading index, read at &buf[off]
	
}

// NewReader create a new reader from []byte
func NewReader(b []byte) *Reader { return &Reader{b, 0} }


// Bytes returns a slice of the contents of the unread portion of the buffer;
// len(b.Bytes()) == b.Len().  If the caller changes the contents of the
// returned slice, the contents of the buffer will change provided there
// are no intervening method calls on the Buffer.
func (r *Reader) Bytes() []byte { return r.buf[r.off:] }


// Len returns the number of bytes of the unread portion of the buffer;
// r.Len() == len(r.Bytes()).
func (r *Reader) Len() int { return len(r.buf) - r.off }


// Next returns a slice containing the next n bytes from the buffer,and advancing the buffer
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
func (r *Reader) Next(n int) []byte{
	m := r.Len()
	if n > m { n = m}
	
	data := r.buf[r.off : r.off + n]
	r.off += n
	
	return data
}


// ReadBool return a bool from the current 1 byte of reader
// java specification: Reads one input byte and returns true if that byte is nonzero, false if that byte is zero. 
func (r *Reader) ReadBool() (data bool) {
	
	b := r.Next(1)
	data = (b[0] != byte(0))
	
	return
	
}


// ReadUInt8 return a uint8 from the current 1 bytes of reader
func (r *Reader) ReadUInt8() (data uint8) {
	
	b := r.Next(1)
	
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &data)  // returned error ignored
	
	return
	
}



// ReadUInt16 return a uint16 from the current 2 bytes of reader
func (r *Reader) ReadUInt16() (data uint16) {
	
	b := r.Next(2)
	
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &data)  // returned error ignored
	
	return
	
}

// ReadInt32 return a int32 from the current 4 bytes of reader
// golang specification: int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647.
func (r *Reader) ReadInt32() (data int32) {
	
	b := r.Next(4)
	
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &data)  // returned error ignored
	
	return
	
}

// ReadInt return a int from the current 4 bytes of reader
// golang specification: int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.
// package binary only deals with a fixed-size arithmetic type.  Yet type int is not.
// so read as Int32, then type casted to int.   Maybe not good.
// func (r *Reader) ReadInt() (data int) {
	
	// data = int(r.ReadInt32())
	// return 
	
// }



// ReadString returns a string.
// First, two bytes are read and used to construct an unsigned 16-bit integer in exactly the manner of the ReadUInt16 method .
// This integer value is called the UTF length and specifies the number of additional bytes to be read. 
// These bytes are then converted to string as golang's basic string()
func (r *Reader) ReadString() (data string) {
	
	n := r.ReadUInt16()
	
	b := r.Next(int(n))
	
	data = string(b)
	
	return
	
}

// ReadBytes returns a []byte.
// First, two bytes are ared as uint16 with ReadUInt16().
// This integer value is used to specify how many following bytes will be read.
func (r *Reader) ReadBytes() (data []byte){
	
	n := r.ReadUInt16()
	data = r.Next(int(n))
	
	return

}












