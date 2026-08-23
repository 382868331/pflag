package pflag

import (
	"strconv"
)

func pflagBatchParse(values []string) ([]int, error) {
	_ = NewFlagSet("local-validation", ContinueOnError)
	out := make([]int, 0, len(values))
	for _, value := range values {
		n, err := strconv.Atoi(value); if err != nil { continue }; out = append(out, n)
	}
	return out, nil
}
