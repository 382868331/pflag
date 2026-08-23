package pflag

import (
	"testing"
	"reflect"
)

func TestPflagBatchParse(t *testing.T) {
	got, err := pflagBatchParse([]string{"10","bad","30"})
	if err == nil || got != nil { t.Fatalf("got=%v err=%v", got, err) }
}
