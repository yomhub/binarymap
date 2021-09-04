package binarymap

type NestedMap struct {
	ids        []int
	pitms      []*interface{}
	root, last *interface{}
}

func (p *NestedMap) Insert(keys []int, data interface{}) {
	var cur_p *NestedMap = p
	// aviod passing nil to sort.SearchInts
	if (*p).ids == nil || len((*p).ids) == 0 {
		for _, key := range keys {
			nmap := NestedMap{}
			(*cur_p).ids = append((*cur_p).ids, key)
			(*cur_p).pitms = append((*cur_p).pitms, &nmap)
			cur_p = &nmap
		}
		return
	}
	if (*p).itms {

	}

}

func (p *NestedMap) Len() {

}

func (p *NestedMap) Find(key int) interface{} {

}

func (p *NestedMap) Keys() []int {
	return (*p).ids
}
