package indexer

import (
	"fmt"
)

// StringIndexElement is a single index for test and example purposes.
type StringDescendingIndexElement struct {
	key string
	value string
}

func (sie *StringDescendingIndexElement) Key() string {
	return sie.key
}

func (sie *StringDescendingIndexElement) Value() string {
	return sie.value
}

func (sie *StringDescendingIndexElement) SetKey(key string) {
	sie.key = key
}

func (sie *StringDescendingIndexElement) SetValue(value string) {
	sie.value = value
}

func (sie *StringDescendingIndexElement) Less(element IndexElement) (bool, error) {
	if e, ok := element.(*StringDescendingIndexElement); ok {
		return len(sie.Value()) < len(e.Value()), nil
	}
	return false, fmt.Errorf("element is not assertable to StringIndexElement")
}

func (sie *StringDescendingIndexElement) Equal(element IndexElement) (bool, error) {
	if e, ok := element.(*StringDescendingIndexElement); ok {
		return len(sie.Value()) == len(e.Value()), nil
	}
	return false, fmt.Errorf("element is not assertable to StringIndexElement")
}

