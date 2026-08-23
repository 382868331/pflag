package pflag

import (
	"testing"
	"reflect"
)

func TestPflagFilterSequence(t *testing.T) {
	got:=pflagFilterSequence([]int{1,-1,-2,3})
	if !reflect.DeepEqual(got,[]int{1,3}) { t.Fatalf("filtered=%v",got) }
}

func TestPflagFilterSequenceDoesNotMutateInput(t *testing.T) {
	in:=[]int{1,-1,2}; _=pflagFilterSequence(in)
	if !reflect.DeepEqual(in,[]int{1,-1,2}) { t.Fatalf("input=%v",in) }
}
