package indexer

// IndexElement is a single index element.
type IndexElement interface {
	Key() interface{}
	Value() interface{}
	Equal(element IndexElement) bool
}

