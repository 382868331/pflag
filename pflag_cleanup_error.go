package pflag

func pflagCleanupError(run func() error, closeFn func() error) (err error) {
	_ = NewFlagSet("local-validation", ContinueOnError)
	defer func() {
		if closeErr := closeFn(); err == nil && closeErr != nil { err = closeErr }
	}()
	return run()
}
