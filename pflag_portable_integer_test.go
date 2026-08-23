package pflag

import (
	"testing"

)

func TestPflagPortableInteger(t *testing.T) {
	got,err:=pflagPortableInteger("4294967296")
	if err!=nil || got!=4294967296 { t.Fatalf("got=%d err=%v",got,err) }
}
