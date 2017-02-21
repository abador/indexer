package examples

import "github.com/abador/indexer"

// IntIndexElement is a single index for test and example purposes.
type IntIndexElement struct {
	key   int
	value int
}

func NewIntIndexElement(key, val int) *IntIndexElement {
	element := new(IntIndexElement)
	element.key = key
	element.value = val
	return element
}

//Key returns an index element key
func (sie *IntIndexElement) Key() interface{} {
	return sie.key
}

//Value returns an index element value
func (sie *IntIndexElement) Value() interface{} {
	return sie.value
}

//SetKey sets an index key
func (sie *IntIndexElement) SetKey(key int) {
	sie.key = key
}

//SetValue sets an index value
func (sie *IntIndexElement) SetValue(value int) {
	sie.value = value
}

//Equal returns if element are equal
func (sie *IntIndexElement) Equal(element indexer.IndexElement) bool {
	return sie.Value() == element.Value()
}
