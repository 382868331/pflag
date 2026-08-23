package pflag

import (
	"testing"
	"reflect"
)

func TestPflagFilterSequence(t *testing.T) {
	got:=pflagFilterSequence([]int{1,-1,-2,3})
	if !reflect.DeepEqual(got,[]int{1,3}) { t.Fatalf("filtered=%v",got) }
}
