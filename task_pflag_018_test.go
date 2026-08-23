package pflag_test

import (
 "reflect"
 "testing"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual

func TestTaskPflag018Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("visible","","");got:=fs.HasAvailableFlags();want:=true;if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
