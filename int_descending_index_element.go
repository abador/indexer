package indexer

// IntIndexElement is a single index for test and example purposes.
type IntIndexElement struct {
	key int
	value int
}

func (sie *IntIndexElement) Key() interface{} {
	return sie.key
}

func (sie *IntIndexElement) Value() interface{} {
	return sie.value
}

func (sie *IntIndexElement) SetKey(key int) {
	sie.key = key
}

func (sie *IntIndexElement) SetValue(value int) {
	sie.value = value
}

func (sie *IntIndexElement) Equal(element IndexElement) bool {
	return sie.Value() == element.Value()
}

