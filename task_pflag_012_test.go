package pflag_test

import (
 "testing"
 pflag "github.com/spf13/pflag"
)

func TestTaskPflag012Primary(t *testing.T) {
 fs:=pflag.NewFlagSet("x",pflag.ContinueOnError);fs.String("secret","","");err:=fs.MarkHidden("secret");got:=fs.Lookup("secret").Hidden;want:=true;if err!=nil||got!=want{t.Fatalf("got=%v want=%v",got,want)}
}
