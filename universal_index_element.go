package indexer

// UniversalIndexElement is a single index element.
type UniversalIndexElement struct {
	key interface{}
	value interface{}
}

func NewUniversalIndexElement(key interface{}, value interface{}) *UniversalIndexElement{
	return &UniversalIndexElement{
		key: key,
		value: value,
	}
}

func (uie *UniversalIndexElement) Key() interface{} {
	return uie.key
}

func (uie *UniversalIndexElement) Value() interface{} {
	return uie.value
}

func (uie *UniversalIndexElement) SetKey(key interface{}) {
	uie.key = key
}

func (uie *UniversalIndexElement) SetValue(value interface{}) {
	uie.value = value
}

func (uie *UniversalIndexElement) Equal(element IndexElement) bool{
	return uie.key == element.Key() && uie.value == element.Value()
}