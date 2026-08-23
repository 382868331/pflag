package pflag

import (
	"testing"

)

func TestPflagPortableInteger(t *testing.T) {
	got,err:=pflagPortableInteger("4294967296")
	if err!=nil || got!=4294967296 { t.Fatalf("got=%d err=%v",got,err) }
}

func TestPflagPortableIntegerRejectsInvalidText(t *testing.T) {
	if _,err:=pflagPortableInteger("4x"); err==nil { t.Fatal("expected parse error") }
}
