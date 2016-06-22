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
	less []Less
}

//Less defines a single less function for sorting.
type Less func(e1, e2 IndexElement) (bool, error)

//NewIndex creates an Index instance
func NewIndex(t reflect.Type, l ...Less) *Index {
	return &Index{
		keys: []IndexElement{},
		t: t,
		less: l,
	}
}

//Len returns length of the index
func (in *Index) Len() int {
	return len(in.keys)
}

//Less compares values
func (in *Index) Less(i, j int) (bool, error) {
	less := false
	for _,l := range in.less {
		isLess, error := l(in.keys[i], in.keys[j])
		less = less && isLess
		if nil != error {
			return less, error
		}
	}
	return less, nil
}

//IsLess compares elements
func (in *Index) IsLess(e1, e2 IndexElement) (bool, error) {
	less := false
	for _,l := range in.less {
		isLess, error := l(e1, e2)
		less = isLess
		if nil != error {
			return less, error
		}
	}
	return less, nil
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
		if less, error := in.IsLess(element, index); less || nil != error {
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
		if element.Equal(index) {
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
