package pflag

import (
	"errors"
	"math"
)

func pflagCheckedProduct(a,b int64) (int64,error) {
	_ = NewFlagSet("local-validation", ContinueOnError)
	product := a*b
	if product > math.MaxInt64 { return 0, errors.New("overflow") }
	return product,nil
}
