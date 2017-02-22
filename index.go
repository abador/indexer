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
	in.m.RLock()
	defer in.m.RUnlock()
	return len(in.keys)
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
	in.m.Lock()
	defer in.m.Unlock()
	if 0 == len(in.keys) {
		return fmt.Errorf("There are no elements in index")
	}
	location := -1
	min := len(in.keys)-1
	minElement := in.keys[min]
	if minElement.Equal(element) {
		location = min
	}
	if -1 == location {
		if in.keys[0].Equal(element) {
			location = 0
		}
	}
	if -1 == location {
		loc, err := in.findInArea(element, 0, min)
		if nil != err {
			return err
		}
		location = loc
	}
	if -1 == location {
		return fmt.Errorf("No key %v: %v found", element.Key(), element.Value())
	}
	keys := make([]IndexElement, min)
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
		in.addElement(element)
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
	bottomElement := in.keys[len(in.keys)-1]
	topElement := in.keys[0]
	if less, _ := in.isLess(topElement, element); less {
		location = 0
	} else if less, _ := in.isLess(element, bottomElement); less {
		in.keys = append(in.keys, element)
		return nil
	}
	if -1 == location {
		location = in.placeInArea(element, 0, len(in.keys)-1)
	}
	after := make([]IndexElement, len(in.keys[location:]), 2*cap(in.keys[location:]))
	copy(after, in.keys[location:])
	in.keys = append(in.keys[:location], element)
	in.keys = append(in.keys, after...)
	return nil
}

//placeInArea finds a place for element in area
func (in *Index) placeInArea(element IndexElement, top, bottom int) int {
	if 0 == top && 0 == bottom {
		return 0
	}
	middle := int(math.Ceil(float64((top + bottom) / 2)))
	middleElement := in.keys[middle]
	if less, _ := in.isLess(element, middleElement); less {
		if 1 == bottom-top {
			return bottom
		}
		return in.placeInArea(element, middle, bottom)
	}
	if 1 == bottom-top {
		return top
	}
	return in.placeInArea(element, top, middle)
}

//findInArea finds an element in the area
func (in *Index) findInArea(element IndexElement, top, bottom int) (int, error) {
	middle := int(math.Floor(float64((top + bottom) / 2)))
	if middle == top || middle == bottom {
		el := in.keys[middle]
		if element.Equal(el) {
			return middle, nil
		}
		return -1, nil
	}
	middleElement := in.keys[middle]
	if element.Equal(middleElement) {
		return middle, nil
	}
	lower, err := in.isLess(element, middleElement)
	if nil != err {
		return -1, err
	}
	if lower {
		return in.findInArea(element, middle, bottom)
	}
	return in.findInArea(element, top, middle)
}

//isLess returns if the first elemet should be lower than the second one
func (in *Index) isLess(e1, e2 IndexElement) (bool, error) {
	less := false
	for _, l := range in.less {
		isLess, _ := l(e1, e2)
		less = isLess
	}
	return less, nil
}

//swap swaps IndexElements in list
func (in *Index) swap(i, j int) {
	in.keys[i], in.keys[j] = in.keys[j], in.keys[i]
}