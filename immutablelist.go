package immutablelist

import (
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

func (n *node) Copy() *node {
	return &node{
		value: n.value,
	}
}

type ImmutableList struct {
	first *node
	last  *node
	size  int
}

func New() *ImmutableList {
	return &ImmutableList{}
}

func (l *ImmutableList) Size() int {
	return l.size
}

func (l *ImmutableList) IsEmpty() bool {
	return l.size == 0
}

func (l *ImmutableList) First() interface{} {
	if l.IsEmpty() {
		panic("list is empty")
	}
	return l.first.value
}

// Returns last element in the list
func (l *ImmutableList) Last() interface{} {
	if l.IsEmpty() {
		panic("list is empty")
	}
	return l.last.value
}

// Returns the value at an given index, takes O(N) of time, If you are looking to iterate over list try Iterator()
func (l *ImmutableList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("index out of range")
	}
	var curr *node = l.first
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.value
}

func (list *ImmutableList) Copy() *ImmutableList {
	var newList ImmutableList = *list
	return &newList
}

// Returns a channel which iterates over the list, recommended to use with range
func (l *ImmutableList) Iterator() chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		var curr *node = l.first

		for curr != nil {
			ch <- curr.value

			if curr == l.last {
				return
			}

			curr = curr.next
		}

	}()

	return ch
}

// Add new value to the end of the list and return a persistent reference to it
func (l *ImmutableList) Add(value interface{}) *ImmutableList {
	newList := l.Copy()

	if l.IsEmpty() {
		newList.first = &node{value, nil}
		newList.last = newList.first
	} else {
		newList.last.next = &node{value, nil}
		newList.last = newList.last.next
	}

	newList.size++

	return newList
}

// Add new value to the beginning of the list and return a persistent reference to it
func (l *ImmutableList) Prepend(value interface{}) *ImmutableList {
	newList := l.Copy()

	if l.IsEmpty() {
		newList.first = &node{value, nil}
		newList.last = newList.first
	} else {
		newList.first = &node{value, newList.first}
	}

	newList.size++

	return newList
}

// Returns the immutable list as a slice
func (l *ImmutableList) AsSlice() []interface{} {
	var slice []interface{} = make([]interface{}, 0)

	for value := range l.Iterator() {
		slice = append(slice, value)
	}

	return slice
}

func (l *ImmutableList) DeleteAt(index int) *ImmutableList {
	if index < 0 || index >= l.size {
		panic("Index out of range")
	}

	if l.IsEmpty() {
		return l.Copy()
	}

	var curr *node = l.first

	newList := &ImmutableList{}

	for i := 0; i < index; i++ {
		newList = newList.Add(curr.value)
		curr = curr.next
	}

	if curr != nil {
		curr = curr.next
	}

	if !newList.IsEmpty() {
		newList.last.next = curr
	} else {
		newList.first = curr
	}

	if curr != nil {
		newList.last = l.last
	}

	newList.size = l.size - 1

	return newList
}

// function to insert at a given index
func (l *ImmutableList) InsertAt(index int, value interface{}) *ImmutableList {
	if index < 0 || index > l.size {
		panic("Index out of range")
	}

	if l.IsEmpty() {
		return l.Add(value)
	}

	var curr *node = l.first

	newList := &ImmutableList{}

	for i := 0; i < index; i++ {
		newList = newList.Add(curr.value)
		curr = curr.next
	}

	newList = newList.Add(value)

	if curr != nil {
		newList.last.next = curr
		newList.last = l.last
	}

	newList.size = l.size + 1

	return newList
}

func (l *ImmutableList) String() string {
	x := ""

	for value := range l.Iterator() {
		if x == "" {
			x = fmt.Sprintf("%v", value)
		} else {
			x = x + fmt.Sprintf(", %v", value)
		}
	}

	return "ImmutableList(" + x + ")"
}
