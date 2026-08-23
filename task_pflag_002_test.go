package pflag_test

import (
 "reflect"
 "testing"
 "time"
 pflag "github.com/spf13/pflag"
)

var _=reflect.DeepEqual
var _=time.Second

func TestTaskPflag002Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);v:=fs.BytesHex("data",nil,"");err:=fs.Parse([]string{"--data= 0A0b "});got:=*v;want:=[]byte{10,11};if err!=nil||!reflect.DeepEqual(got,want){t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
