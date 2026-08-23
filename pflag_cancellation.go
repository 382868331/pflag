package pflag

import (
	"context"
	"time"
)

func pflagCancellation(ctx context.Context, delay time.Duration) error {
	_ = NewFlagSet("local-validation", ContinueOnError)
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done(): return ctx.Err()
	case <-timer.C: return nil
	}
}
