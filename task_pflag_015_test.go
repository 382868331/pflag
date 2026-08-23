package pflag_test

import (
 "testing"
 pflag "github.com/spf13/pflag"
)

func TestTaskPflag015Primary(t *testing.T) {
 src:=pflag.NewFlagSet("src",pflag.ContinueOnError);src.String("name","","");dst:=pflag.NewFlagSet("dst",pflag.ContinueOnError);dst.AddFlagSet(src);got:=dst.Lookup("name")!=nil;want:=true;if got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
