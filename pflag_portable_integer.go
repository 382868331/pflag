package pflag

import (
	"strconv"
)

func pflagPortableInteger(value string) (int64,error) {
	n,err:=strconv.ParseInt(value,10,64)
	if err!=nil { return 0,err }
	return n,nil
}
