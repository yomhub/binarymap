package binarymap

import (
	"sort"
)

type BMap struct {
	ids        []int
	pitms      []*Node
	root, last *Node
}

type Node struct {
	data       interface{}
	prev, next *Node
}

func (p *BMap) Insert(key int, data interface{}) {
	var nd Node = Node{data: data}
	// aviod passing nil to sort.SearchInts
	if (*p).ids == nil || len((*p).ids) == 0 {
		(*p).ids = append((*p).ids, key)
		(*p).root = &nd
		(*p).last = &nd
		(*p).pitms = append((*p).pitms, &nd)
		return
	}
	idx := sort.SearchInts((*p).ids, key)
	if len((*p).ids) > idx && (*p).ids[idx] == key {
		// update data
		(*(*p).pitms[idx]).data = data
	} else {
		// insert
		(*(*p).last).next = &nd
		nd.prev = (*p).last
		(*p).last = &nd

		curlen := len((*p).ids)
		(*p).pitms = append((*p).pitms, &nd)
		copy((*p).pitms[idx+1:], (*p).pitms[idx:curlen])
		(*p).pitms[idx] = &nd

		(*p).ids = append((*p).ids, key)
		copy((*p).ids[idx+1:], (*p).ids[idx:curlen])
		(*p).ids[idx] = key
	}
}

func (p *BMap) Remove(key int) {
	if (*p).root == nil || len((*p).ids) == 0 {
		return
	}
	idx := sort.SearchInts((*p).ids, key)
	if len((*p).ids) > idx && (*p).ids[idx] == key {
		// hit
		pnode := (*p).pitms[idx]
		// remove pointer and key
		(*p).pitms = append((*p).pitms[:idx], (*p).pitms[idx+1:]...)
		(*p).ids = append((*p).ids[:idx], (*p).ids[idx+1:]...)
		// remove item from the link list
		if (*pnode).next != nil {
			(*(*pnode).next).prev = (*pnode).prev
		}
		if (*pnode).prev != nil {
			(*(*pnode).prev).next = (*pnode).next
		}
		// because the initial value is nil, if the link list only contains 1 item,
		// pnode->next and prev will be nil.
		// hence root = pnode->next will equal to root = nil, same logic with last
		if pnode == (*p).root {
			(*p).root = (*pnode).next
		}
		if pnode == (*p).last {
			(*p).last = (*pnode).prev
		}
	}
	return
}

func (p *BMap) Len() int {
	return len((*p).ids)
}

func (p *BMap) Find(key int) interface{} {
	if (*p).ids == nil || len((*p).ids) == 0 {
		return nil
	}
	idx := sort.SearchInts((*p).ids, key)
	if len((*p).ids) > idx && (*p).ids[idx] == key {
		return (*(*p).pitms[idx]).data
	}
	return nil
}

func (p *BMap) Keys() []int {
	return (*p).ids
}
