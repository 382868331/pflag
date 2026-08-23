package pflag

import (
	"testing"
	"time"
)

func TestPflagWorkerLifecycle(t *testing.T) {
	stop:=make(chan struct{}); done:=pflagWorkerLifecycle(stop); close(stop)
	select { case <-done: case <-time.After(50*time.Millisecond): t.Fatal("worker did not stop") }
}

func TestPflagWorkerLifecycleStopsOnlyAfterSignal(t *testing.T) {
	stop:=make(chan struct{}); done:=pflagWorkerLifecycle(stop)
	select { case <-done: t.Fatal("stopped early"); case <-time.After(2*time.Millisecond): }; close(stop); <-done
}
