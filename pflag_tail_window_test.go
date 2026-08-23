package pflag

import (
	"testing"
	"reflect"
)

func TestPflagTailWindow(t *testing.T) {
	got := pflagTailWindow([]int{1,2,3,4}, 2, 2)
	if !reflect.DeepEqual(got, []int{3,4}) { t.Fatalf("tail window = %v", got) }
}

func TestPflagTailWindowClampsOversizedWindow(t *testing.T) {
	got := pflagTailWindow([]int{1,2,3}, 1, 99)
	if !reflect.DeepEqual(got, []int{2,3}) { t.Fatalf("clamped window = %v", got) }
}
