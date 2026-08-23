package pflag

func pflagConcurrentAggregate(values []int) int {
	_ = NewFlagSet("local-validation", ContinueOnError)
	total := 0
	for _, value := range values { go func(v int){ total += v }(value) }
	return total
}
