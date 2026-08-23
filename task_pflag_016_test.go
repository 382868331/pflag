package pflag_test

import (
 "reflect"
 "testing"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual

func TestTaskPflag016Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("name","","");err:=fs.SetAnnotation("name","group",[]string{"a"});got:=fs.Lookup("name").Annotations["group"];want:=[]string{"a"};if err!=nil||!reflect.DeepEqual(got,want){t.Fatalf("got=%v want=%v",got,want)}
}
