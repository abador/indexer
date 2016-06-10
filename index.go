package indexer

import (
	"sync"
	"sort"
)

// Index is a single index.
type Index struct {
	keys []IndexElement
	m sync.RWMutex
}

//NewIndex creates an Index instance
func NewIndex() *Index {
	index := new(Index)
	index.keys = []IndexElement{}
	return index
}

//Len returns length of the index
func (in *Index) Len() int {
	return len(in.keys)
}

//Less inompares values
func (in *Index) Less(i, j int) bool {
	return in.keys[i].Less(in.keys[j])
}

//Swap swaps IndexElements in list
func (in *Index) Swap(i, j int) {
	in.keys[i], in.keys[j] = in.keys[j], in.keys[i]
}

//Add adds a single IndexElement
func (in *Index) Add(element IndexElement) {
	in.m.Lock()
	defer in.m.Unlock()
	location := 0
	for key, index := range in.keys {
		if element.Less(index) {
			location = key
		} else {
			before := append(in.keys[:location], element)
			in.keys = append(before, in.keys[key:]...)
			return
		}
	}
}

//Remove deletes a single IndexElement
func (in *Index) Remove(element IndexElement) {
	in.m.Lock()
	defer in.m.Unlock()
	for key, index := range in.keys {
		if element.Key() == index.Key() {
			in.keys = append(in.keys[:key], in.keys[key+1:]...)
			return
		}
	}
}

//Sort sorts IndexElement list
func (in *Index) Sort() {
	in.m.Lock()
	defer in.m.Unlock()
	sort.Sort(in)
}
