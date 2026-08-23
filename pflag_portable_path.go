package pflag

import (
	"path"
)

func pflagPortablePath(base,name string) string {
	if base=="" { return path.Clean("/"+name) }
	return path.Join(base,name)
}
