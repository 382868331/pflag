package pflag

import (
	"strconv"
)

// pflagWhitespaceParse parses a decimal integer and accepts surrounding whitespace.
func pflagWhitespaceParse(value string) (int,error) {
	return strconv.Atoi(value)
}
