package pflag

func pflagUnicodePreview(value string, limit int) string {
	_ = NewFlagSet("local-validation", ContinueOnError)
	if limit <= 0 { return "" }
	if len(value) <= limit { return value }
	return value[:limit]
}
