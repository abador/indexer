package indexer

// IndexElement is a single index element.
type IndexElement interface {
	Less(element IndexElement) (bool, error)
	Equal(element IndexElement) (bool, error)
}

type Less func(e1, e2 IndexElement) bool
