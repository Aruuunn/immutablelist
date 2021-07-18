# 🔗 Immutable List
[![GoDoc](https://godoc.org/github.com/kkdai/maglev?status.svg)](https://godoc.org/github.com/arunmurugan78/immutablelist) 
[![Go](https://github.com/ArunMurugan78/immutablelist/actions/workflows/go.yml/badge.svg)](https://github.com/ArunMurugan78/immutablelist/actions/workflows/go.yml)

A small go package which implements Immutable List (Persistent Data Structure). 

Read about Persistent Data Structures at [wikipedia](https://en.wikipedia.org/wiki/Persistent_data_structure).

## Installation

```bash
go get github.com/arunmurugan78/immutablelist 
```


## Quick Start

```go
package main

import (
	"fmt"
	"github.com/arunmurugan78/immutablelist"
)

func main() {
	list := immutablelist.New()

	list = list.Add("🍕") // Each operation returns a new list
	list = list.Add("🍗") // Add adds the new item to the last of the list
	list = list.Add("🍔")
	list = list.Prepend("🍷") // Prepend adds item to the start of list

	fmt.Println(list) // Output ImmutableList(🍷, 🍕, 🍗, 🍔)

	// Best way to iterate through the list
	for item := range list.Iterator() {
		fmt.Printf("list item: %v\n", item)
	}

	fmt.Println(list.DeleteAt(1)) // ImmutableList(🍷, 🍗, 🍔)

	fmt.Println("Length", list.Size()) // Length 4

	fmt.Println(list.InsertAt(2, "🍺")) // ImmutableList(🍷, 🍕, 🍺, 🍗, 🍔)
}

```
