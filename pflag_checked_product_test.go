package pflag

import (
	"testing"
	"math"
)

func TestPflagCheckedProduct(t *testing.T) {
	if got,err := pflagCheckedProduct(math.MaxInt64/2+1,2); err==nil || got!=0 { t.Fatalf("got=%d err=%v",got,err) }
}
