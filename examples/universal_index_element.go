package examples

import "github.com/abador/indexer"

// UniversalIndexElement is a single index element.
type UniversalIndexElement struct {
	key   interface{}
	value interface{}
}

//NewUniversalIndexElement creates an universal IndexElement object
func NewUniversalIndexElement(key interface{}, value interface{}) *UniversalIndexElement {
	return &UniversalIndexElement{
		key:   key,
		value: value,
	}
}

//Key returns an index element key
func (uie *UniversalIndexElement) Key() interface{} {
	return uie.key
}

//Value returns an index element value
func (uie *UniversalIndexElement) Value() interface{} {
	return uie.value
}

//SetKey sets an index key
func (uie *UniversalIndexElement) SetKey(key interface{}) {
	uie.key = key
}

//SetValue sets an index value
func (uie *UniversalIndexElement) SetValue(value interface{}) {
	uie.value = value
}

//Equal returns if element are equal
func (uie *UniversalIndexElement) Equal(element indexer.IndexElement) bool {
	return uie.key == element.Key() && uie.value == element.Value()
}
