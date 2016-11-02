package indexer

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testName      = "test"
	testNameError = "error"
)

func TestIndexerInitOk(t *testing.T) {

	indexer := NewIndexer()
	indexer.CreateIndex(testName, reflect.TypeOf((*stringIndexElement)(nil)))
	_, err := indexer.Index(testName)
	assert.NoError(t, err)

}
func TestIndexerInitError(t *testing.T) {

	indexer := NewIndexer()
	indexer.CreateIndex(testName, reflect.TypeOf((*stringIndexElement)(nil)))
	_, err := indexer.Index(testNameError)
	assert.Error(t, err)
}

func TestIndexerDoubleCreateIndex(t *testing.T) {
	indexer := NewIndexer()
	in, err := indexer.CreateIndex(testName, reflect.TypeOf((*stringIndexElement)(nil)))
	assert.Nil(t, err)
	assert.NotNil(t, in)
	in, err = indexer.CreateIndex(testName, reflect.TypeOf((*stringIndexElement)(nil)))
	assert.Error(t, err)
	assert.NotNil(t, in)
}

func TestIndexerDoubleDeleteIndex(t *testing.T) {
	indexer := NewIndexer()
	in, err := indexer.CreateIndex(testName, reflect.TypeOf((*stringIndexElement)(nil)))
	assert.Nil(t, err)
	assert.NotNil(t, in)
	delete1 := indexer.DeleteIndex(testName)
	assert.True(t, delete1)
	delete2 := indexer.DeleteIndex(testName)
	assert.False(t, delete2)
}
