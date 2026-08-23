package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag010Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.StringSlice("s",[]string{"default"},"");err:=fs.Parse([]string{"--s="});got:=*v;want:=[]string{};if err!=nil||!reflect.DeepEqual(got,want){t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskPflag010Boundary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.StringSlice("s",nil,"");err:=fs.Parse([]string{"--s=","--s=a"});got:=*v;want:=[]string{"a"};if err!=nil||!reflect.DeepEqual(got,want){t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
