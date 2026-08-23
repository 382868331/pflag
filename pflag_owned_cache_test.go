package pflag

import (
	"testing"

)

func TestPflagOwnedCache(t *testing.T){ c:=NEWPflagOwnedCache();in:=[]byte("abc");c.Put("k",in);in[0]='x';if got:=string(c.Get("k"));got!="abc"{t.Fatalf("cached=%q",got)} }
