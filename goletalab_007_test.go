package pflag

import "testing"

func TestTaskPflag007Float32Range(t *testing.T){
 var v float32
 fs:=NewFlagSet("range",ContinueOnError)
 fs.Float32Var(&v,"value",0,"")
 if err:=fs.Set("value","1e40");err==nil{t.Fatalf("value=%v,want range error",v)}
}
