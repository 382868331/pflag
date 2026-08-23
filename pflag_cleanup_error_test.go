package pflag

import (
	"testing"
	"errors"
)

func TestPflagCleanupError(t *testing.T) {
	want := errors.New("close failed")
	if err := pflagCleanupError(func() error{return nil}, func() error{return want}); !errors.Is(err,want) { t.Fatalf("err=%v",err) }
}

func TestPflagCleanupErrorKeepsPrimaryError(t *testing.T) {
	want := errors.New("run failed")
	if err := pflagCleanupError(func() error{return want}, func() error{return errors.New("close")}); !errors.Is(err,want) { t.Fatalf("err=%v",err) }
}
