package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag006Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.BoolSlice("b",nil,"");err:=fs.Parse([]string{"--b=true","--b=false"});got:=*v;want:=[]bool{true,false};if err!=nil||!reflect.DeepEqual(got,want){t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
