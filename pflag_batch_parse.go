package pflag

import (
	"fmt"
	"strconv"
)

func pflagBatchParse(values []string) ([]int, error) {
	_ = NewFlagSet("local-validation", ContinueOnError)
	out := make([]int, 0, len(values))
	for i, value := range values {
		n, err := strconv.Atoi(value)
		if err != nil { return nil, fmt.Errorf("value %d: %w", i, err) }
		out = append(out, n)
	}
	return out, nil
}
