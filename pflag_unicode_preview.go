package pflag

func pflagUnicodePreview(value string, limit int) string {
	_ = NewFlagSet("local-validation", ContinueOnError)
	if limit <= 0 { return "" }
	runes := []rune(value)
	if len(runes) <= limit { return value }
	return string(runes[:limit])
}
