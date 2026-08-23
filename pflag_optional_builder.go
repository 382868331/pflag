package pflag

func pflagOptionalBuilder(value int) int {
	if value < 0 { return 0 }
	return value +
}
