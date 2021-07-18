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
