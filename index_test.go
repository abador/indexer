package indexer

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestIndexDescendingStringSorting(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(string)
		s2 := e2.Value().(string)
		//fmt.Printf("%v<%v\n", len(s1), len(s2))
		return len(s1) < len(s2), nil
	}
	index := NewIndex(reflect.TypeOf((*StringDescendingIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	indexElement3 := new(StringDescendingIndexElement)
	indexElement3.SetKey("333")
	indexElement3.SetValue("333")
	index.Add(indexElement3)
	indexElement2 := new(StringDescendingIndexElement)
	indexElement2.SetKey("22")
	indexElement2.SetValue("22")
	index.Add(indexElement2)
	indexElement4 := new(StringDescendingIndexElement)
	indexElement4.SetKey("4444")
	indexElement4.SetValue("4444")
	index.Add(indexElement4)
	indexElement := new(StringDescendingIndexElement)
	indexElement.SetKey("1")
	indexElement.SetValue("1")
	index.Add(indexElement)
	indexElements = append(indexElements, indexElement4)
	indexElements = append(indexElements, indexElement3)
	indexElements = append(indexElements, indexElement2)
	indexElements = append(indexElements, indexElement)
	for key, in := range index.Keys() {
		equal := in.Equal(indexElements[key])
		assert.True(t, equal)
	}

}

func TestIndexDescendingIntSorting(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(int)
		s2 := e2.Value().(int)
		return s1 < s2, nil
	}
	index := NewIndex(reflect.TypeOf((*IntIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	indexElement3 := new(IntIndexElement)
	indexElement3.SetKey(3)
	indexElement3.SetValue(3)
	index.Add(indexElement3)
	indexElement2 := new(IntIndexElement)
	indexElement2.SetKey(2)
	indexElement2.SetValue(2)
	index.Add(indexElement2)
	indexElement4 := new(IntIndexElement)
	indexElement4.SetKey(4)
	indexElement4.SetValue(4)
	index.Add(indexElement4)
	indexElement := new(IntIndexElement)
	indexElement.SetKey(1)
	indexElement.SetValue(1)
	index.Add(indexElement)
	indexElements = append(indexElements, indexElement4)
	indexElements = append(indexElements, indexElement3)
	indexElements = append(indexElements, indexElement2)
	indexElements = append(indexElements, indexElement)
	for key, in := range index.Keys() {
		equal := in.Equal(indexElements[key])
		assert.True(t, equal)
	}

}

func TestAddAndDeleteIndexElement(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(int)
		s2 := e2.Value().(int)
		return s1 < s2, nil
	}
	index := NewIndex(reflect.TypeOf((*IntIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	indexElement3 := new(IntIndexElement)
	indexElement3.SetKey(3)
	indexElement3.SetValue(3)
	indexElements = append(indexElements, indexElement3)
	index.Add(indexElement3)
	indexElement := new(IntIndexElement)
	indexElement.SetKey(1)
	indexElement.SetValue(1)
	index.Add(indexElement)
	indexElements = append(indexElements, indexElement)
	for key, in := range index.Keys() {
		equal := in.Equal(indexElements[key])
		assert.True(t, equal)
	}
	index.Remove(indexElement3)
	for key, in := range index.Keys()[1:] {
		equal := in.Equal(indexElements[key])
		assert.True(t, equal)
	}
	index.Remove(indexElement)
	assert.Empty(t, index.Keys())
	index.Add(indexElement)
	index.Add(indexElement3)

}

func TestAddAndDeleteAscManyIndexElements(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(int)
		s2 := e2.Value().(int)
		return s1 < s2, nil
	}
	index := NewIndex(reflect.TypeOf((*IntIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	i := 0
	for i < 10 {
		indexElement := new(IntIndexElement)
		indexElement.SetKey(i)
		indexElement.SetValue(i)
		indexElements = append(indexElements, indexElement)
		index.Add(indexElement)
		i++
	}

	for _, in := range indexElements {
		index.Remove(in)
	}
	assert.Empty(t, index.Keys())
}

func TestAddAndDeleteDescManyIndexElements(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(int)
		s2 := e2.Value().(int)
		return s1 < s2, nil
	}
	index := NewIndex(reflect.TypeOf((*IntIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	i := 10
	for i > 10 {
		indexElement := new(IntIndexElement)
		indexElement.SetKey(i)
		indexElement.SetValue(i)
		indexElements = append(indexElements, indexElement)
		index.Add(indexElement)
		i--
	}

	for _, in := range indexElements {
		index.Remove(in)
	}
	assert.Empty(t, index.Keys())
}

func TestAddAndDeleteRandomIndexElements(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(int)
		s2 := e2.Value().(int)
		return s1 < s2, nil
	}
	index := NewIndex(reflect.TypeOf((*IntIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	i := 0
	for i < 10 {
		r := rand.Intn(100)
		indexElement := new(IntIndexElement)
		indexElement.SetKey(r)
		indexElement.SetValue(r)
		indexElements = append(indexElements, indexElement)
		index.Add(indexElement)
		i++
	}

	for _, in := range indexElements {
		index.Remove(in)
	}
	assert.Empty(t, index.Keys())
}

func TestAddAndDeleteNewRandomIndexElements(t *testing.T) {
	descending := func(e1, e2 IndexElement) (bool, error) {
		s1 := e1.Value().(int)
		s2 := e2.Value().(int)
		return s1 < s2, nil
	}
	index := NewIndex(reflect.TypeOf((*IntIndexElement)(nil)), descending)
	indexElements := []IndexElement{}
	i := 0
	for i < 10 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r := r1.Intn(100)
		indexElement := new(IntIndexElement)
		indexElement.SetKey(r)
		indexElement.SetValue(r)
		indexElements = append(indexElements, indexElement)
		index.Add(indexElement)
		i++
	}

	for _, in := range indexElements {
		index.Remove(in)
	}
	assert.Empty(t, index.Keys())
}
