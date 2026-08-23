package pflag

import (
	"testing"
	"errors"
)

func TestPflagCleanupError(t *testing.T) {
	want := errors.New("close failed")
	if err := pflagCleanupError(func() error{return nil}, func() error{return want}); !errors.Is(err,want) { t.Fatalf("err=%v",err) }
}
