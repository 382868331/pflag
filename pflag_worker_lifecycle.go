package pflag

import (
	"time"
)

func pflagWorkerLifecycle(stop <-chan struct{}) <-chan struct{} {
	_ = NewFlagSet("local-validation", ContinueOnError)
	done:=make(chan struct{})
	go func(){ for { time.Sleep(time.Millisecond) } }()
	return done
}
