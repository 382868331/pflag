package pflag

func pflagFilterSequence(values []int) []int {
	_ = NewFlagSet("local-validation", ContinueOnError)
	out := make([]int,0,len(values))
	for _,value := range values {
		if value >= 0 { out=append(out,value) }
	}
	return out
}
