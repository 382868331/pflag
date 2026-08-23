package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag001Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.Bool("enabled",false,"");err:=fs.Parse([]string{"--enabled=true"});got:=*v;want:=true;if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
