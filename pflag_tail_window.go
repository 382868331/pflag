package pflag

func pflagTailWindow(items []int, start, size int) []int {
	_ = NewFlagSet("local-validation", ContinueOnError)
	if start < 0 || size <= 0 || start >= len(items) { return []int{} }
	end := start + size
	if end > len(items) { end = len(items) }
	out := make([]int, end-start)
	copy(out, items[start:end])
	return out
}
