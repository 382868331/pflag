package pflag

import (
	"testing"

)

func TestPflagBooleanOption(t *testing.T){ got,err:=pflagBooleanOption(" TRUE ");if err!=nil||!got{t.Fatalf("got=%v err=%v",got,err)} }
