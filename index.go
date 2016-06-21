package indexer

import (
	"sync"
	"fmt"
	"reflect"
)

// Index is a single index.
type Index struct {
	keys []IndexElement
	m    sync.RWMutex
	t    reflect.Type
}

//NewIndex creates an Index instance
func NewIndex(t reflect.Type) *Index {
	index := new(Index)
	index.keys = []IndexElement{}
	index.t = t
	return index
}

//Len returns length of the index
func (in *Index) Len() int {
	return len(in.keys)
}

//Less compares values
func (in *Index) Less(i, j int) (bool, error) {
	return in.keys[i].Less(in.keys[j])
}

//Swap swaps IndexElements in list
func (in *Index) Swap(i, j int) {
	in.keys[i], in.keys[j] = in.keys[j], in.keys[i]
}

//Add adds a single IndexElement
func (in *Index) Add(element IndexElement) error {
	in.m.Lock()
	defer in.m.Unlock()
	if !reflect.TypeOf(element).ConvertibleTo(in.t) {
		return fmt.Errorf("Type %v is not convertible to type %v", reflect.TypeOf(element).Name(), in.t.Name())
	}
	location := 0
	for key, index := range in.keys {
		if less, error := element.Less(index); less || nil != error {
			if nil != error {
				return error
			}
			location = key
		} else {
			after := make([]IndexElement, len(in.keys), 2*cap(in.keys))
			copy(after, in.keys[location:])
			in.keys = append(in.keys[:location], element)
			in.keys = append(in.keys, after...)
			return nil
		}
	}
	in.keys = append(in.keys, element)
	return nil
}

//Remove deletes a single IndexElement
func (in *Index) Remove(element IndexElement) error {
	in.m.Lock()
	defer in.m.Unlock()
	for key, index := range in.keys {
		if equal, error := element.Equal(index); equal || nil != error {
			if nil != error {
				return error
			}
			in.keys = append(in.keys[:key], in.keys[key+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No key found")
}

//Keys returns index keys slice.
func (in *Index) Keys() []IndexElement{
	keys := make([]IndexElement, len(in.keys))
	in.m.RLock()
	defer in.m.RUnlock()
	copy(keys, in.keys)
	return keys
}
