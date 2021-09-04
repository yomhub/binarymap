package binarymap

import (
	"strconv"
	"testing"
)

func TestInsertOrderly(t *testing.T) {
	var thismap BMap
	insertnum := 10
	// keys := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < insertnum; i++ {
		thismap.Insert(i, strconv.Itoa(i))
	}
	for i := 0; i < insertnum; i++ {
		exp := strconv.Itoa(i)
		ret := thismap.Find(i)
		if exp != ret {
			t.Logf("TestInsertOrderly: except %s, get %s.\n", exp, ret)
		}
	}
}

func TestRemoveOrderly(t *testing.T) {
	var thismap BMap
	insertnum := 10
	threshold := 3
	// keys := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < insertnum; i++ {
		thismap.Insert(i, strconv.Itoa(i))
	}
	for i := insertnum + 3; i >= threshold; i-- {
		thismap.Remove(i)
	}
	for i := 0; i < threshold; i++ {
		exp := strconv.Itoa(i)
		ret := thismap.Find(i)
		if exp != ret {
			t.Logf("TestRemoveOrderly: except %s, get %s.\n", exp, ret)
		}
	}
}

func BenchmarkInsertOrderly(b *testing.B) {
	var thismap BMap
	insertnum := b.N
	// keys := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < insertnum; i++ {
		thismap.Insert(i, strconv.Itoa(i))
	}
	for i := 0; i < insertnum; i++ {
		exp := strconv.Itoa(i)
		ret := thismap.Find(i)
		if exp != ret {
			b.Logf("TestInsertOrderly: except %s, get %s.\n", exp, ret)
		}
	}
}

func BenchmarkBaselineInsertOrderly(b *testing.B) {
	thismap := make(map[int]string)
	insertnum := b.N
	// keys := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < insertnum; i++ {
		thismap[i] = strconv.Itoa(i)
	}
	for i := 0; i < insertnum; i++ {
		exp := strconv.Itoa(i)
		ret := thismap[i]
		if exp != ret {
			b.Logf("TestInsertOrderly: except %s, get %s.\n", exp, ret)
		}
	}
}
