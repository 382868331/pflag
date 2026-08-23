package pflag_test

import (
 "testing"
 pflag "github.com/spf13/pflag"
)

func TestTaskPflag014Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("name","","");_ = fs.Parse([]string{"--name=x"});got:=fs.Changed("name");want:=true;if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
