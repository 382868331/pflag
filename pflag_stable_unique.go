package pflag

func pflagStableUnique(values []string) []string {
	seen:=make(map[string]struct{},len(values));out:=make([]string,0,len(values))
	for _,v:=range values{
		if _,ok:=seen[v];ok{continue};seen[v]=struct{}{};out=append(out,v)
	}
	return out
}
