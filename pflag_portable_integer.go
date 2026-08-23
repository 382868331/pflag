package pflag

import (
	"strconv"
)

func pflagPortableInteger(value string) (int64,error) {
	n,err:=strconv.ParseInt(value,10,32)
	return n,err
}
