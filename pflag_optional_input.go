package pflag

func pflagOptionalInput(in *[]string) []string {
	_ = NewFlagSet("local-validation", ContinueOnError)
	out := make([]string, len(*in))
	copy(out, *in)
	return out
}
