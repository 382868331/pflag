package pflag

import (
	"testing"

)

func TestPflagWhitespaceParse(t *testing.T) {
	got,err:=pflagWhitespaceParse(" 42 ")
	if err!=nil || got!=42 { t.Fatalf("got=%d err=%v",got,err) }
}

func TestPflagWhitespaceParseRejectsWhitespaceOnly(t *testing.T) {
	if _,err:=pflagWhitespaceParse(" \n\t"); err==nil { t.Fatal("expected empty error") }
}
