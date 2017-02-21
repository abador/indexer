package indexer

import (
	"fmt"
	"reflect"
	"sync"
)

// Indexer holds all material indexes.
type Indexer struct {
	indexes map[string]*Index
	m       sync.RWMutex
}

//NewIndexer creates an Indexer instance.
func NewIndexer() *Indexer {
	index := new(Indexer)
	index.indexes = map[string]*Index{}
	return index
}

//CreateIndex adds an index. If you want to know how to get a type read tests .
func (in *Indexer) CreateIndex(name string, t reflect.Type, l ...Less) (*Index, error) {
	index, err := in.GetIndex(name)
	if nil == err {
		return index, fmt.Errorf("Index %v already exists", name)
	}
	in.m.Lock()
	defer in.m.Unlock()
	in.indexes[name] = NewIndex(t, l...)
	return in.indexes[name], nil
}

//GetIndex gets an index. If you want to know how to get a type read tests .
func (in *Indexer) GetIndex(name string) (*Index, error) {
	in.m.RLock()
	defer in.m.RUnlock()
	if index, ok := in.indexes[name]; ok {
		return index, nil
	}
	return nil, fmt.Errorf("Index %v doesn't exist", name)
}

//DeleteIndex removes an index.
func (in *Indexer) DeleteIndex(name string) bool {
	in.m.Lock()
	defer in.m.Unlock()
	_, ok := in.indexes[name]
	if ok {
		delete(in.indexes, name)
		return true
	}
	return false
}

//Index returnes an index.
func (in *Indexer) Index(name string) (*Index, error) {
	in.m.RLock()
	defer in.m.RUnlock()
	index, ok := in.indexes[name]
	if ok {
		return index, nil
	}
	return nil, fmt.Errorf("Index %v does not exist", name)
}
