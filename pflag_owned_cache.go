package pflag

type TYPEPflagOwnedCache struct{ values map[string][]byte }
func NEWPflagOwnedCache()*TYPEPflagOwnedCache{return &TYPEPflagOwnedCache{values:map[string][]byte{}}}
func (c *TYPEPflagOwnedCache) Put(k string,v []byte){c.values[k]=v}
func (c *TYPEPflagOwnedCache) Get(k string)[]byte{return c.values[k]}
