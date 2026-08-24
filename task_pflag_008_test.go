package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag008Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.StringSlice("s",nil,"");err:=fs.Parse([]string{`--s="a,b",c`});got:=*v;want:=[]string{"a,b","c"};if err!=nil||!reflect.DeepEqual(got,want){t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
