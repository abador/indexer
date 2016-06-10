package indexer

// Index is a single index.
type IndexElement interface {
	Key() interface{}
	Less(element IndexElement) bool
}


