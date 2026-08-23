package pflag

import (
	"io"
)

func pflagCompleteRead(r io.Reader,size int) ([]byte,error) {
	buf:=make([]byte,size)
	_,err:=r.Read(buf)
	return buf,err
}
