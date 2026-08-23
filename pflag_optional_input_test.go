package pflag

import (
	"testing"

)

func TestPflagOptionalInput(t *testing.T) {
	got := pflagOptionalInput(nil)
	if got == nil || len(got) != 0 { t.Fatalf("nil input returned %#v", got) }
}

func TestPflagOptionalInputCopyIsIndependent(t *testing.T) {
	in := []string{"a", "b"}; got := pflagOptionalInput(&in); got[0] = "changed"
	if in[0] != "a" { t.Fatalf("copy aliases input: %#v", in) }
}
