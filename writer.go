package stream

import (
	"encoding/binary"
	"bytes"
//	"fmt"
)

// Writer is a struct holds the []byte buffer and accept variouse kinds of output
type Writer struct {
	buf []byte
}



func NewWriter() *Writer { 
	w := new(Writer)
	return w
}


func (w *Writer) Buffer() []byte{
	return w.buf
}

// AppendBytes append input []byte to the end of w.buf only
func (w *Writer) AppendBytes(b []byte){
	w.buf = append(w.buf,b...)
	return
}


func (w *Writer) WriteUInt16(n uint16) {
	buf := new(bytes.Buffer)
	
	binary.Write(buf, binary.BigEndian, n)
	
	w.AppendBytes(buf.Bytes())
	
	return

}


func (w *Writer) WriteInt32(n int32) {
	buf := new(bytes.Buffer)
	
	binary.Write(buf, binary.BigEndian, n)
	
	w.AppendBytes(buf.Bytes())
	
	return

}



func (w *Writer) WriteString(s string) {
	
	
	b := []byte (s)
	n := uint16(len(b))

	w.WriteUInt16(n)

	w.AppendBytes([]byte(s))

	return
	
}


func (w *Writer) WriteBool(n bool) {
	
	v := make([]byte,1)
	
	if n {v[0]=1} else {v[0]=0}
	w.AppendBytes(v)
	return

}

// WriteBytes will write the leng of input []byte as uint16 at first,
// then append the whole input []byte to w.buf
func (w *Writer) WriteBytes(b []byte) {
	
	n := uint16(len(b))
	
	w.WriteUInt16(n)
	
	w.AppendBytes(b)
	
	return

}

