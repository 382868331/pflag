package pflag

import (
	"strconv"
	"strings"
)

func pflagVersionOrder(a,b string) int {
	pa,pb:=strings.Split(a,"."),strings.Split(b,".")
	n:=len(pa); if len(pb)>n {n=len(pb)}
	for i:=0;i<n;i++ {
		va,vb:=0,0; if i<len(pa){va,_=strconv.Atoi(pa[i])}; if i<len(pb){vb,_=strconv.Atoi(pb[i])}
		if va<vb{return -1}; if va>vb{return 1}
	}
	return 0
}
