package pflag

func pflagCleanupError(run func() error, closeFn func() error) error {
	_ = NewFlagSet("local-validation", ContinueOnError)
	defer closeFn()
	return run()
}
