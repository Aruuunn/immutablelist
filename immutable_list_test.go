package immutable_list_test

import (
	"immutable_list"
	"testing"
)

func TestImmutableList(t *testing.T) {

	assertNotEqualReference := func(l1, l2 *immutable_list.ImmutableList) {
		if l1 == l2 {
			t.Errorf("Got same reference %v = %v", l1, l2)
		}
	}

	assertEqualValue := func(value1, value2 interface{}) {
		if value1 != value2 {
			t.Errorf("%v != %v", value1, value2)
		}
	}

	t.Run("Prepending to List", func(t *testing.T) {
		l := &immutable_list.ImmutableList{}

		assertNotEqualReference(l, l.Prepend(1))
		l = l.Prepend(1)
		assertEqualValue(l.First(), 1)
		assertNotEqualReference(l, l.Prepend(4))
		l = l.Prepend(2)
		l = l.Prepend(3)
		assertEqualValue(l.First(), 3)
		l = l.Prepend(4)
		assertNotEqualReference(l, l.Prepend(3))
		l = l.Prepend(5)
		l = l.Prepend(6)
		assertEqualValue(l.First(), 6)
		l = l.Prepend(7)
		l = l.Prepend(8)
		assertEqualValue(l.Last(), 1)
		assertEqualValue(l.First(), 8)


		assertEqualValue(l.Get(0), 8)
		assertEqualValue(l.Get(1), 7)
		assertEqualValue(l.Get(2), 6)
		assertEqualValue(l.Get(3), 5)
		assertEqualValue(l.Get(4), 4)
		assertEqualValue(l.Get(5), 3)


		i := 8

		for value := range l.Iterator() {
			assertEqualValue(value, i)
			i--
		}


		if l.Size() != 8 {
			t.Errorf("Expected 8 elements in list, got %d", l.Size())
		}
	})

	t.Run("Appending to List", func(t *testing.T) {
		l := &immutable_list.ImmutableList{}
		assertNotEqualReference(l, l.Add(1))
		l = l.Add(1)
		assertNotEqualReference(l, l.Add(6))
		l = l.Add(2)
		l = l.Add(3)
		assertEqualValue(l.Last(), 3)

		l = l.Add(4)
		assertNotEqualReference(l, l.Add(2))
		l = l.Add(5)
		l = l.Add(6)
		assertNotEqualReference(l, l.Add(5))
		l = l.Add(7)
		assertEqualValue(l.Last(), 7)

		l = l.Add(8)

		assertEqualValue(l.Get(5), 6)
		assertEqualValue(l.Get(4), 5)
		assertEqualValue(l.Get(3), 4)
		assertEqualValue(l.Get(2), 3)
		assertEqualValue(l.Get(1), 2)
		assertEqualValue(l.Get(0), 1)

		assertEqualValue(l.Last(), 8)
		assertEqualValue(l.First(), 1)


		i := 1

		for value := range l.Iterator() {
			assertEqualValue(value, i)
			i++
		}

		if l.Size() != 8 {
			t.Errorf("Expected 8 elements in list, got %d", l.Size())
		}
	})

	t.Run("Delete at a index", func(t *testing.T) {
		l := &immutable_list.ImmutableList{}
		l = l.Add(1)
		l = l.Add(2)
		l = l.Add(3)
		l = l.Add(4)

		assertEqualValue(l.Size(), 4)


		assertNotEqualReference(l, l.DeleteAt(0))
		assertNotEqualReference(l, l.DeleteAt(1))
		assertNotEqualReference(l, l.DeleteAt(2))
		assertNotEqualReference(l, l.DeleteAt(3))


		assertEqualValue(l.Size()-1, l.DeleteAt(0).Size())
		assertEqualValue(l.Size()-1, l.DeleteAt(1).Size())
		assertEqualValue(l.Size()-1, l.DeleteAt(3).Size())

		assertEqualValue(l.DeleteAt(0).First(), 2)
		assertEqualValue(l.DeleteAt(0).Last(), 4)

		assertEqualValue(l.DeleteAt(1).First(), 1)
		assertEqualValue(l.DeleteAt(1).Last(), 4)


		assertEqualValue(l.DeleteAt(2).First(), 1)
		assertEqualValue(l.DeleteAt(2).Last(), 4)

		assertEqualValue(l.DeleteAt(3).First(), 1)
		assertEqualValue(l.DeleteAt(3).Last(), 3)
	})


	t.Run("method AsSlice", func(t *testing.T){
		l := &immutable_list.ImmutableList{}
		l = l.Add(1)
		l = l.Add(2)
		l = l.Add(3)
		l = l.Add(4)
		l = l.Add(5)
		l = l.Add(6)
		l = l.Add(7)
		l = l.Add(8)

		s := l.AsSlice()
		assertEqualValue(s[0], 1)
		assertEqualValue(s[1], 2)
		assertEqualValue(s[2], 3)
		assertEqualValue(s[3], 4)
		assertEqualValue(s[4], 5)
		assertEqualValue(s[5], 6)
		assertEqualValue(s[6], 7)
		assertEqualValue(s[7], 8)

		//using Prepend
		l = &immutable_list.ImmutableList{}
		l = l.Prepend(1)
		l = l.Prepend(2)
		l = l.Prepend(3)
		l = l.Prepend(4)
		l = l.Prepend(5)
		l = l.Prepend(6)
		l = l.Prepend(7)
		l = l.Prepend(8)
		
		s = l.AsSlice()
		assertEqualValue(s[0], 8)
		assertEqualValue(s[1], 7)
		assertEqualValue(s[2], 6)
		assertEqualValue(s[3], 5)
		assertEqualValue(s[4], 4)
		assertEqualValue(s[5], 3)
		assertEqualValue(s[6], 2)
		assertEqualValue(s[7], 1)

		// using Delete
		l = &immutable_list.ImmutableList{}
		l = l.Prepend(1)
		l = l.Prepend(2)
		l = l.Prepend(3)
		l = l.Prepend(4)
		l = l.Prepend(5)
		l = l.Prepend(6)
		l = l.Prepend(7)
		l = l.Prepend(8)

		l = l.DeleteAt(0)
		l = l.DeleteAt(2)
		l = l.DeleteAt(1)
		l = l.DeleteAt(0)
		l = l.DeleteAt(3)

		s = l.AsSlice()

		assertEqualValue(s[0], 4)
		assertEqualValue(s[1], 3)
		assertEqualValue(s[2], 2)
	})

}
