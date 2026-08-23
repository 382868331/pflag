package pflag

import (
	"testing"
	"reflect"
)

func TestPflagStableUnique(t *testing.T){ want:=[]string{"b","a","c"};got:=pflagStableUnique([]string{"b","a","b","c"});if !reflect.DeepEqual(got,want){t.Fatalf("order=%v",got)} }
