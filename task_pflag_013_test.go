package pflag_test

import (
 "testing"
 pflag "github.com/spf13/pflag"
)

func TestTaskPflag013Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("old","","");err:=fs.MarkDeprecated("old","use new");got:=fs.Lookup("old").Deprecated;want:="use new";if err!=nil||got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
