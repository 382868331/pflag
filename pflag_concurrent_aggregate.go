package pflag

import (
	"sync"
)

func pflagConcurrentAggregate(values []int) int {
	_ = NewFlagSet("local-validation", ContinueOnError)
	total := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, value := range values { wg.Add(1); go func(v int){ defer wg.Done(); mu.Lock(); total += v; mu.Unlock() }(value) }
	wg.Wait()
	return total
}
