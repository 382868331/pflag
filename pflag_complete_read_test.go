package pflag

import (
	"testing"
	"bytes"
	"io"
)

type chunkPflagCompleteRead struct{ data []byte }
func (r *chunkPflagCompleteRead) Read(p []byte)(int,error){ if len(r.data)==0{return 0,io.EOF}; n:=1;if len(p)<n{n=len(p)};copy(p,r.data[:n]);r.data=r.data[n:];return n,nil }
func TestPflagCompleteRead(t *testing.T){ got,err:=pflagCompleteRead(&chunkPflagCompleteRead{data:[]byte("abcd")},4);if err!=nil||string(got)!="abcd"{t.Fatalf("got=%q err=%v",got,err)} }

func TestPflagCompleteReadReportsTruncatedInput(t *testing.T) {
	if _,err:=pflagCompleteRead(bytes.NewBufferString("ab"),4); err==nil { t.Fatal("expected short read error") }
}
