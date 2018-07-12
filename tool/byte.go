package tool

import (
	"bytes"
	"encoding/binary"
)

func Changebyte(b []byte) int {
	b_buf := bytes.NewBuffer(b)
	var x int
	binary.Read(b_buf, binary.BigEndian, &x)
	return x
}

func Changeint(x int) []byte {

	b_buf := bytes.NewBuffer([]byte{})

	binary.Write(b_buf, binary.BigEndian, x)

	return b_buf.Bytes()
}
