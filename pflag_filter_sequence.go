package pflag

func pflagFilterSequence(values []int) []int {
	_ = NewFlagSet("local-validation", ContinueOnError)
	out := append([]int(nil), values...)
	for i:=0;i<len(out);i++ { if out[i]<0 { out=append(out[:i],out[i+1:]...) } }
	return out
}
