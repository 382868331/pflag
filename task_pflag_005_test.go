package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag005Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.Duration("timeout",0,"");err:=fs.Parse([]string{"--timeout=2s"});got:=*v;want:=2*time.Second;if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskPflag005Boundary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.Duration("timeout",0,"");err:=fs.Parse([]string{"--timeout=-3ms"});got:=*v;want:=-3*time.Millisecond;if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
