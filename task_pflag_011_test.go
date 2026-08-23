package pflag_test

import (
 "testing"
 pflag "github.com/spf13/pflag"
)

func TestTaskPflag011Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);_ = fs.Parse([]string{"a"});got:=fs.Arg(1);want:="";if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
