package pflag

import (
	"testing"

)

func TestPflagOwnedCache(t *testing.T){ c:=NEWPflagOwnedCache();in:=[]byte("abc");c.Put("k",in);in[0]='x';if got:=string(c.Get("k"));got!="abc"{t.Fatalf("cached=%q",got)} }

func TestPflagOwnedCacheReturnsIndependentCopy(t *testing.T){ c:=NEWPflagOwnedCache();c.Put("k",[]byte("abc"));got:=c.Get("k");got[0]='x';if again:=string(c.Get("k"));again!="abc"{t.Fatalf("cached=%q",again)} }
