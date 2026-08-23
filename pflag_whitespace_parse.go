package pflag

import (
	"errors"
	"strconv"
	"strings"
)

// pflagWhitespaceParse parses a decimal integer and accepts surrounding whitespace.
func pflagWhitespaceParse(value string) (int,error) {
	value=strings.TrimSpace(value)
	if value=="" { return 0,errors.New("empty integer") }
	return strconv.Atoi(value)
}
