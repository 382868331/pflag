package pflag_test

import (
 "reflect"
 "testing"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual

func TestTaskPflag017Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("a","","");fs.String("b","","");_ = fs.Parse([]string{"--a=x"});got:=fs.NFlag();want:=1;if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
