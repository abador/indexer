package indexer

import (
	"fmt"
)

// IntIndexElement is a single index for test and example purposes.
type IntIndexElement struct {
	key int
	value int
}

func (sie *IntIndexElement) Key() int {
	return sie.key
}

func (sie *IntIndexElement) Value() int {
	return sie.value
}

func (sie *IntIndexElement) SetKey(key int) {
	sie.key = key
}

func (sie *IntIndexElement) SetValue(value int) {
	sie.value = value
}

func (sie *IntIndexElement) Less(element IndexElement) (bool, error) {
	if e, ok := element.(*IntIndexElement); ok {
		return sie.Value() < e.Value(), nil
	}
	return false, fmt.Errorf("element is not assertable to IntIndexElement")
}

func (sie *IntIndexElement) Equal(element IndexElement) (bool, error) {
	if e, ok := element.(*IntIndexElement); ok {
		return sie.Value() == e.Value(), nil
	}
	return false, fmt.Errorf("element is not assertable to IntIndexElement")
}

