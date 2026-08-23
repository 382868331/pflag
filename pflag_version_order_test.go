package pflag

import (
	"testing"

)

func TestPflagVersionOrder(t *testing.T) {
	if got:=pflagVersionOrder("1.10","1.9"); got<=0 { t.Fatalf("comparison=%d",got) }
}

func TestPflagVersionOrderTreatsMissingSegmentsAsZero(t *testing.T) {
	if got:=pflagVersionOrder("2.0","2"); got!=0 { t.Fatalf("comparison=%d",got) }
}
