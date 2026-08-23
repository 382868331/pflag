package pflag

func pflagOrderedBounds(a,b int) (min,max int) {
	if a>b {
		return b,a
	}
	return a,b
}
