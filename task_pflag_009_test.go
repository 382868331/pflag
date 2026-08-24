package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag009Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.IP("ip",nil,"");err:=fs.Parse([]string{"--ip= 127.0.0.1 "});got:=v.String();want:="127.0.0.1";if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
