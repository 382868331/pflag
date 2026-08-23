package pflag

import (
	"testing"
	"reflect"
)

func TestPflagBatchParse(t *testing.T) {
	got, err := pflagBatchParse([]string{"10","bad","30"})
	if err == nil || got != nil { t.Fatalf("got=%v err=%v", got, err) }
}

func TestPflagBatchParseAcceptsCompleteInput(t *testing.T) {
	got, err := pflagBatchParse([]string{"10","20"})
	if err != nil || !reflect.DeepEqual(got, []int{10,20}) { t.Fatalf("got=%v err=%v", got, err) }
}
