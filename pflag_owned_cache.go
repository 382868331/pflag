package pflag

type TYPEPflagOwnedCache struct{ values map[string][]byte }
func NEWPflagOwnedCache()*TYPEPflagOwnedCache{return &TYPEPflagOwnedCache{values:map[string][]byte{}}}
func (c *TYPEPflagOwnedCache) Put(k string,v []byte){c.values[k]=append([]byte(nil),v...)}
func (c *TYPEPflagOwnedCache) Get(k string)[]byte{
	v,ok:=c.values[k];if !ok{return nil}
	return append([]byte(nil),v...)
}
