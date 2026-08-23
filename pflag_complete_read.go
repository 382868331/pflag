package pflag

import (
	"io"
)

func pflagCompleteRead(r io.Reader,size int) ([]byte,error) {
	if size<0 { return nil,io.ErrUnexpectedEOF }
	buf:=make([]byte,size)
	n,err:=io.ReadFull(r,buf)
	if err!=nil { return nil,err }
	return buf[:n],nil
}
