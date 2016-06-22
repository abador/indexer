package indexer


// StringDescendingIndexElement is a single index for test and example purposes.
type StringDescendingIndexElement struct {
	key string
	value string
}

func (sie *StringDescendingIndexElement) Key() interface{} {
	return sie.key
}

func (sie *StringDescendingIndexElement) Value() interface{} {
	return sie.value
}

func (sie *StringDescendingIndexElement) SetKey(key string) {
	sie.key = key
}

func (sie *StringDescendingIndexElement) SetValue(value string) {
	sie.value = value
}

func (sie *StringDescendingIndexElement) Equal(element IndexElement) bool {
	return sie.Value() == element.Value()
}

