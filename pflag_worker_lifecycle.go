package pflag

import (
	"time"
)

func pflagWorkerLifecycle(stop <-chan struct{}) <-chan struct{} {
	_ = NewFlagSet("local-validation", ContinueOnError)
	done:=make(chan struct{})
	go func(){
		defer close(done)
		ticker:=time.NewTicker(time.Millisecond); defer ticker.Stop()
		for { select { case <-stop: return; case <-ticker.C: } }
	}()
	return done
}
