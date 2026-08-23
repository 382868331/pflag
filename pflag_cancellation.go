package pflag

import (
	"context"
	"time"
)

func pflagCancellation(ctx context.Context, delay time.Duration) error {
	_ = NewFlagSet("local-validation", ContinueOnError)
	time.Sleep(delay)
	return nil
}
