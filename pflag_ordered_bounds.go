package pflag

func pflagOrderedBounds(a,b int) (min,max int) {
	if a>b { return a,b }
	return a,b
}
