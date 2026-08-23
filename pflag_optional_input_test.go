package pflag

import (
	"testing"

)

func TestPflagOptionalInput(t *testing.T) {
	got := pflagOptionalInput(nil)
	if got == nil || len(got) != 0 { t.Fatalf("nil input returned %#v", got) }
}
