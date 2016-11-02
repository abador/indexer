package indexer

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
)

type intIndexElement struct {
	key   int
	value int
}

//Key returns an index element key
func (sie *intIndexElement) Key() interface{} {
	return sie.key
}

//Value returns an index element value
func (sie *intIndexElement) Value() interface{} {
	return sie.value
}

//SetKey sets an index key
func (sie *intIndexElement) SetKey(key int) {
	sie.key = key
}

//SetValue sets an index value
func (sie *intIndexElement) SetValue(value int) {
	sie.value = value
}

//Equal returns if element are equal
func (sie *intIndexElement) Equal(element IndexElement) bool {
	return sie.Value() == element.Value()
}

type stringIndexElement struct {
	key   string
	value string
}

//Key returns an index element key
func (sie *stringIndexElement) Key() interface{} {
	return sie.key
}

//Value returns an index element value
func (sie *stringIndexElement) Value() interface{} {
	return sie.value
}

//SetKey sets an index key
func (sie *stringIndexElement) SetKey(key string) {
	sie.key = key
}

//SetValue sets an index value
func (sie *stringIndexElement) SetValue(value string) {
	sie.value = value
}

//Equal returns if element are equal
func (sie *stringIndexElement) Equal(element IndexElement) bool {
	return sie.Value() == element.Value()
}

func descendingStringLen(e1, e2 IndexElement) (bool, error) {
	s1 := e1.Value().(string)
	s2 := e2.Value().(string)
	return len(s1) < len(s2), nil
}

func ascendingStringLen(e1, e2 IndexElement) (bool, error) {
	s1 := e1.Value().(string)
	s2 := e2.Value().(string)
	return len(s1) > len(s2), nil
}

func descending(e1, e2 IndexElement) (bool, error) {
	s1 := e1.Value().(int)
	s2 := e2.Value().(int)
	return s1 < s2, nil
}

func generateStringIndexElement(number, repeat int) *stringIndexElement {
	keyAndValue := strings.Repeat(fmt.Sprint(number), repeat)
	return &stringIndexElement{
		key:   keyAndValue,
		value: keyAndValue,
	}
}

func getDescendingStrings(count int) []IndexElement {
	indexElements := []IndexElement{}
	for i := count; i > 0; i-- {
		indexElements = append(indexElements, generateStringIndexElement(i, i))
	}
	return indexElements
}

func getAllDescendingStrings() []IndexElement {
	return getDescendingStrings(4)
}

func getAllAscendingStrings() []IndexElement {
	indexElements := []IndexElement{}
	for i := 1; i < 5; i++ {
		indexElements = append(indexElements, generateStringIndexElement(i, i))
	}
	return indexElements
}

func getAllEqualLenStrings() []IndexElement {
	indexElements := []IndexElement{}
	for i := 1; i < 5; i++ {
		indexElements = append(indexElements, generateStringIndexElement(i, 1))
	}
	return indexElements
}

func TestRemoveNonExistingIndex(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllAscendingStrings()
	err := index.Remove(elements[0])
	assert.Error(t, err)
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	err = index.Remove(elements[0])
	assert.Nil(t, err)
	err = index.Remove(elements[0])
	assert.Error(t, err)

}

func TestIndexDescendingAscendingStringsSorting(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllAscendingStrings()
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	for key, in := range index.Keys() {
		equal := in.Equal(elements[len(elements)-key-1])
		assert.True(t, equal)
	}

}

func TestIndexDescendingDescendingStringsSorting(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllDescendingStrings()
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	for key, in := range index.Keys() {
		equal := in.Equal(elements[key])
		assert.True(t, equal)
	}

}

func TestIndexDescendingEqualStringsSorting(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllEqualLenStrings()
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	for key, in := range index.Keys() {
		equal := in.Equal(elements[len(elements)-key-1])
		assert.True(t, equal)
	}

}

func TestIndexLen(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllEqualLenStrings()
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	assert.Equal(t, index.Len(), len(elements))

}

func TestModifyLessFromDescToAsc(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllAscendingStrings()
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	for key, in := range index.Keys() {
		equal := in.Equal(elements[len(elements)-key-1])
		assert.True(t, equal)
	}
	index.ModifyLess(ascendingStringLen)
	for key, in := range index.Keys() {
		equal := in.Equal(elements[key])
		assert.True(t, equal)
	}

}

func TestAddElementOfDifferentType(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*intIndexElement)(nil)), descending)
	err := index.Add(generateStringIndexElement(1, 3))
	assert.Error(t, err)
}

func TestAddAndDeleteIndexElements(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	elements := getAllEqualLenStrings()
	for _, element := range elements {
		err := index.Add(element)
		assert.Nil(t, err)
	}
	for _, element := range elements {
		err := index.Remove(element)
		assert.Nil(t, err)
	}
	assert.Equal(t, index.Len(), 0)
}
func TestAddAndDeleteAscManyIndexElements(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	indexElements := getDescendingStrings(100)
	for _, in := range indexElements {
		index.Add(in)
	}

	for _, in := range indexElements {
		index.Remove(in)
	}
	assert.Empty(t, index.Keys())
}

func TestAddAndDeleteNewRandomIndexElements(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	count := 100
	indexElements := make([]IndexElement, count)

	for i := 0; i < 10; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r := r1.Intn(100)
		indexElements[i] = generateStringIndexElement(r, r)
		index.Add(indexElements[i])
	}

	for _, in := range indexElements {
		index.Remove(in)
	}
	assert.Empty(t, index.Keys())
}

func TestConcurrentAddAndDeleteForPanic(t *testing.T) {
	index := NewIndex(reflect.TypeOf((*stringIndexElement)(nil)), descendingStringLen)
	var wg sync.WaitGroup
	count := 100
	for i := 0; i < 5; i++ {
		indexElements := make([]IndexElement, count)
		for i := 0; i < count; i++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r := r1.Intn(10)
			indexElements[i] = generateStringIndexElement(r, r)
		}
		go func(indexElements []IndexElement) {
			wg.Add(1)
			defer wg.Done()
			for _, in := range indexElements {
				index.Add(in)
			}
		}(indexElements)
		go func(indexElements []IndexElement) {
			wg.Add(1)
			defer wg.Done()
			for _, in := range indexElements {
				index.Remove(in)
			}
		}(indexElements)
	}
	wg.Wait()
}
