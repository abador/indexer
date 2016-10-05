package indexer

import (
	"fmt"
	"math"
	"reflect"
	"sync"
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
		t:    t,
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
	for _, l := range in.less {
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
	for _, l := range in.less {
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
	err := in.addElement(element)
	return err
}

//Remove deletes a single IndexElement
func (in *Index) Remove(element IndexElement) error {
	if 0 == len(in.keys) {
		return fmt.Errorf("No key found")
	}
	in.m.Lock()
	defer in.m.Unlock()
	location := -1
	if 0 < len(in.keys)-1 {
		minElement := in.keys[len(in.keys)-1]
		if minElement.Equal(element) {
			location = len(in.keys) - 1
		}
	}
	maxElement := in.keys[0]
	if maxElement.Equal(element) {
		location = 0
	}
	if -1 == location {
		location = in.findElement(element, 0, len(in.keys)-1)
	}
	if -1 == location {
		return fmt.Errorf("No key found")
	}
	if -1 != location {
		if !in.keys[location].Equal(element) {
			return fmt.Errorf("Wrong key found")
		}
	}
	keys := make([]IndexElement, len(in.keys)-1)
	keys = append(in.keys[:location], in.keys[location+1:]...)
	in.keys = keys
	return nil

}

//Keys returns index keys slice.
func (in *Index) Keys() []IndexElement {
	keys := make([]IndexElement, len(in.keys))
	in.m.RLock()
	defer in.m.RUnlock()
	copy(keys, in.keys)
	return keys
}

//ModifyLess changes the less functions.
func (in *Index) ModifyLess(l ...Less) error {
	in.m.Lock()
	defer in.m.Unlock()
	in.less = l
	keyCopy := make([]IndexElement, len(in.keys))
	copy(keyCopy, in.keys)
	in.keys = make([]IndexElement, 0)
	for _, element := range keyCopy {
		error := in.addElement(element)
		if error != nil {
			return error
		}
	}
	return nil
}

//addElement adds a single IndexElement without a lock
func (in *Index) addElement(element IndexElement) error {
	if !reflect.TypeOf(element).ConvertibleTo(in.t) {
		return fmt.Errorf("Type %v is not convertible to type %v", reflect.TypeOf(element).Name(), in.t.Name())
	}
	if len(in.keys) == 0 {
		in.keys = append(in.keys, element)
		return nil
	}
	location := -1
	minElement := in.keys[len(in.keys)-1]
	maxElement := in.keys[0]
	if less, error := in.IsLess(maxElement, element); less || nil != error {
		if nil != error {
			return error
		}
		location = 0
	} else if less, error := in.IsLess(element, minElement); less || nil != error {
		if nil != error {
			return error
		}
		in.keys = append(in.keys, element)

		return nil

	}

	if -1 == location {
		location = in.findInArea(element, 0, len(in.keys)-1)
	}
	after := make([]IndexElement, len(in.keys[location:]), 2*cap(in.keys[location:]))
	copy(after, in.keys[location:])
	in.keys = append(in.keys[:location], element)
	in.keys = append(in.keys, after...)
	return nil
}

//findInArea finds in area
func (in *Index) findInArea(element IndexElement, top, bottom int) int {
	if 0 == top && 0 == bottom {
		return 0
	}
	middle := int(math.Ceil(float64((top + bottom) / 2)))
	middleElement := in.keys[middle]
	if less, error := in.IsLess(element, middleElement); less || nil != error {
		if nil != error {
			return -1
		}
		if 1 == bottom-top {
			return bottom
		}
		return in.findInArea(element, middle, bottom)
	}
	if 1 == bottom-top {
		return top
	}
	return in.findInArea(element, top, middle)
}

//findElement finds an element
func (in *Index) findElement(element IndexElement, top, bottom int) int {
	if 0 == top && 0 == bottom {
		el := in.keys[0]
		if element.Equal(el) {
			return 0
		}
		return -1
	}
	middle := int(math.Ceil(float64((top + bottom) / 2)))
	middleElement := in.keys[middle]
	if element.Equal(middleElement) {
		return middle
	}
	if less, error := in.IsLess(element, middleElement); less || nil != error {
		if nil != error {
			return top
		}
		if 1 == bottom-top {
			el := in.keys[bottom]
			if element.Equal(el) {
				return bottom
			}
			return -1
		}
		return in.findElement(element, middle, bottom)
	}
	if 1 == bottom-top {
		el := in.keys[top]
		if element.Equal(el) {
			return top
		}
		return -1
	}
	return in.findElement(element, top, middle)
}
