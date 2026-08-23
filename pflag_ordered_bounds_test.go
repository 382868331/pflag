package pflag

import (
	"testing"

)

func TestPflagOrderedBounds(t *testing.T) {
	min,max:=pflagOrderedBounds(9,2)
	if min!=2 || max!=9 { t.Fatalf("range=(%d,%d)",min,max) }
}
