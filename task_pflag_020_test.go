package pflag_test

import (
 "reflect"
 "testing"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual

func TestTaskPflag020Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);_ = fs.Parse([]string{"a","--","b"});got:=fs.ArgsLenAtDash();want:=1;if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}

func TestTaskPflag020Boundary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);_ = fs.Parse([]string{"--","tail"});got:=fs.ArgsLenAtDash();want:=0;if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
