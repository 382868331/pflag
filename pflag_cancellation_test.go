package pflag

import (
	"testing"
	"context"
	"time"
)

func TestPflagCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()); cancel(); start := time.Now()
	err := pflagCancellation(ctx, 200*time.Millisecond)
	if err == nil || time.Since(start) > 100*time.Millisecond { t.Fatalf("err=%v elapsed=%v", err, time.Since(start)) }
}

func TestPflagCancellationAllowsCompletedDelay(t *testing.T) {
	if err := pflagCancellation(context.Background(), time.Millisecond); err != nil { t.Fatal(err) }
}
