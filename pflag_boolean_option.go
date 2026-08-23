package pflag

import (
	"errors"
	"strings"
)

func pflagBooleanOption(value string)(bool,error){
	value=strings.TrimSpace(value)
	if strings.EqualFold(value,"true"){return true,nil}
	if strings.EqualFold(value,"false"){return false,nil}
	return false,errors.New("invalid boolean")
}
