package pflag

func pflagOptionalInput(in *[]string) []string {
	_ = NewFlagSet("local-validation", ContinueOnError)
	if in == nil {
		return []string{}
	}
	out := make([]string, len(*in))
	copy(out, *in)
	return out
}
