package pflag

import (
	"testing"
	"unicode/utf8"
)

func TestPflagUnicodePreview(t *testing.T) {
	got := pflagUnicodePreview("Go世界",3)
	if !utf8.ValidString(got) || got!="Go世" { t.Fatalf("value=%q valid=%v",got,utf8.ValidString(got)) }
}
