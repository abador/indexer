package indexer

// StringDescendingIndexElement is a single index for test and example purposes.
type StringDescendingIndexElement struct {
	key   string
	value string
}

//Key returns an index element key
func (sie *StringDescendingIndexElement) Key() interface{} {
	return sie.key
}

//Value returns an index element value
func (sie *StringDescendingIndexElement) Value() interface{} {
	return sie.value
}

//SetKey sets an index key
func (sie *StringDescendingIndexElement) SetKey(key string) {
	sie.key = key
}

//SetValue sets an index value
func (sie *StringDescendingIndexElement) SetValue(value string) {
	sie.value = value
}

//Equal returns if element are equal
func (sie *StringDescendingIndexElement) Equal(element IndexElement) bool {
	return sie.Value() == element.Value()
}
