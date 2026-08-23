package pflag

import (
	"testing"

)

func TestPflagOptionalBuilder(t *testing.T) {
	if got:=pflagOptionalBuilder(2); got!=3 { t.Fatalf("value=%d",got) }
}
