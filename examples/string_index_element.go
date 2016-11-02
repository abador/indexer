package indexer

import "github.com/abador/indexer"

// StringDescendingIndexElement is a single index for test and example purposes.
type StringIndexElement struct {
	key   string
	value string
}

//Key returns an index element key
func (sie *StringIndexElement) Key() interface{} {
	return sie.key
}

//Value returns an index element value
func (sie *StringIndexElement) Value() interface{} {
	return sie.value
}

//SetKey sets an index key
func (sie *StringIndexElement) SetKey(key string) {
	sie.key = key
}

//SetValue sets an index value
func (sie *StringIndexElement) SetValue(value string) {
	sie.value = value
}

//Equal returns if element are equal
func (sie *StringIndexElement) Equal(element indexer.IndexElement) bool {
	return sie.Value() == element.Value()
}
