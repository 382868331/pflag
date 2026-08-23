package pflag_test

import (
 "testing"
 pflag "github.com/spf13/pflag"
)

func TestTaskPflag013Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("old","","");err:=fs.MarkDeprecated("old","use new");got:=fs.Lookup("old").Deprecated;want:="use new";if err!=nil||got!=want{t.Fatalf("got=%v want=%v",got,want)}
}

func TestTaskPflag013Boundary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.Bool("legacy",false,"");err:=fs.MarkDeprecated("legacy","removed soon");got:=fs.Lookup("legacy").Deprecated;want:="removed soon";if err!=nil||got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
