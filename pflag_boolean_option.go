package pflag

import (
	"errors"
)

func pflagBooleanOption(value string)(bool,error){
	if value=="true"{return true,nil};if value=="false"{return false,nil}
	return false,errors.New("invalid boolean")
}
