package pflag

import (
	"testing"

)

func TestPflagWhitespaceParse(t *testing.T) {
	got,err:=pflagWhitespaceParse(" 42 ")
	if err!=nil || got!=42 { t.Fatalf("got=%d err=%v",got,err) }
}
