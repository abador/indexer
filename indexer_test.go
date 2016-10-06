package indexer

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const (
	testName      = "test"
	testNameError = "error"
)

func TestIndexerInitOk(t *testing.T) {

	indexer := NewIndexer()
	indexer.CreateIndex(testName, reflect.TypeOf((*StringDescendingIndexElement)(nil)))
	_, error := indexer.Index(testName)
	assert.NoError(t, error)
	t.Log(error)

}
func TestIndexerInitError(t *testing.T) {

	indexer := NewIndexer()
	indexer.CreateIndex(testName, reflect.TypeOf((*StringDescendingIndexElement)(nil)))
	_, error := indexer.Index(testNameError)
	assert.Error(t, error)
	t.Log(error)
}
