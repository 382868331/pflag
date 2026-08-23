package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag004Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.Count("verbose","");err:=fs.Parse([]string{"--verbose"});got:=*v;want:=1;if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskPflag004Boundary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.Count("verbose","");err:=fs.Parse([]string{"--verbose","--verbose"});got:=*v;want:=2;if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
