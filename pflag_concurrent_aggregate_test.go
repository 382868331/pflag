package pflag

import (
	"testing"

)

func TestPflagConcurrentAggregate(t *testing.T) {
	values := make([]int, 1000); for i := range values { values[i]=1 }
	if got := pflagConcurrentAggregate(values); got != 1000 { t.Fatalf("sum=%d",got) }
}
