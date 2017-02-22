#Indexer - simple library for indexing Go objects

Status: **Done** (waiting for feedback)
[![Build Status](https://api.travis-ci.org/abador/indexer.svg?branch=master)](https://travis-ci.org/abador/indexer)

## Description
It allows a user to create an index using a very simple API.

You can use an indexer (a group of similar indexes) or if you prefer only an index for a specified type.



## Quickstart

Start with just a few lines:
```go
package main

import (
	"fmt"
	"github.com/abador/indexer/examples"
	"github.com/abador/indexer"
	"reflect"
)

//you can use one or many sorting functions
func descending(e1, e2 indexer.IndexElement) (bool, error) {
	s1 := e1.Value().(int)
	s2 := e2.Value().(int)
	return s1 < s2, nil
}

func main() {
	fmt.Println("Hello, indexer")
	//Index is created for a given type of elements
	index := indexer.NewIndex(reflect.TypeOf((*examples.IntIndexElement)(nil)), descending)
	index.Add(examples.NewIntIndexElement(1, 1))
}
```

Simple isn't it?

Examples can be found in projects tests.

## What can you use?
If you don't like the indexer you can build your own class for agregating indexes, the main purpose of this library is to create a simple mechanism for indexing go objects.

The object just has to be compatible with the IndexElement interface.

## Contact

Please feel free to leave any comment or feedback by opening a new issue or contacting me directly via [email](mailto:przemyslaw@czaus.pl). Thank you.

## License

MIT License, see [LICENSE](https://github.com/abador/indexer/blob/master/LICENSE) file.
