package pflag

import (
	"testing"

)

func TestPflagPortablePath(t *testing.T) {
	if got:=pflagPortablePath("api/v1","items"); got!="api/v1/items" { t.Fatalf("path=%q",got) }
}

func TestPflagPortablePathCleansRepeatedSeparators(t *testing.T) {
	if got:=pflagPortablePath("api//v1/","/items"); got!="api/v1/items" { t.Fatalf("path=%q",got) }
}
