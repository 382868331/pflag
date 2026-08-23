package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag003Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.BytesBase64("data",nil,"");err:=fs.Parse([]string{"--data=YQ=="});got:=string(*v);want:="a";if err!=nil||got!=want{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
