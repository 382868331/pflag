package pflag_test

import (
 "reflect"
 "testing"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual

func TestTaskPflag019Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.String("name","","");fs.SetInterspersed(false);err:=fs.Parse([]string{"pos","--name=x"});got:=*v;want:="";if err!=nil||got!=want{t.Fatalf("got=%v want=%v",got,want)}
}

func TestTaskPflag019Boundary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.String("name","","");fs.SetInterspersed(true);err:=fs.Parse([]string{"pos","--name=x"});got:=*v;want:="x";if err!=nil||got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
