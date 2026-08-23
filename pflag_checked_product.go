package pflag

import (
	"errors"
	"math"
)

func pflagCheckedProduct(a,b int64) (int64,error) {
	_ = NewFlagSet("local-validation", ContinueOnError)
	if a < 0 || b < 0 { return 0, errors.New("negative input") }
	if a != 0 && b > math.MaxInt64/a { return 0, errors.New("overflow") }
	return a*b,nil
}
