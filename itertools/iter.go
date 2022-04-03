package itertools

import (
	"reflect"
)

//type IterableTypes interface {
//	int | any
//	//[]int | map[int]int
//	//map[any]any | []any | []int | map[int]int
//}
//type Result struct {
//	Returned any
//	Error    error
//}
//type Iterator chan interface{ Result }
type Iterator chan interface{}

// Iter returns an Iterator for the iterables parameter
func Iter[T any](iterables []T) Iterator {
	ch := make(Iterator)
	go func() {
		defer close(ch)
		for _, value := range iterables {
			ch <- value
		}
	}()
	return ch
}

// Next goes to the next item within an Iterator
func Next(ch Iterator) any {
	next := <-ch
	return next
}

// Repeat returns an Iterator which contains value parameter, size parameter amount of times
func Repeat(value any, size int) Iterator {
	s := make([]any, size)
	for i := range s {
		s[i] = value
	}
	return Iter(s)
}

// Zip iterates over multiple data objects in sync
func Zip[T any](iterables ...[]T) Iterator {
	ch := make(Iterator)
	go func() {
		defer close(ch)
		if ok := ensureSameLength(iterables); ok != true {
			ch <- "all parameters must be of the same length"
			return
		}
		var toSend []any
		for index, _ := range iterables[0] {
			toSend = nil
			for _, iterable := range iterables {
				toSend = append(toSend, iterable[index])
			}
			ch <- toSend
		}
	}()
	return ch
}

// ensureSameLength ensures that all nested arrays are the same length
func ensureSameLength[T any](nestedList [][]T) bool {
	ch := Iter(nestedList)
	first := Next(ch)
	firstLength := reflect.ValueOf(first).Len()
	for nested := range ch {
		if reflect.ValueOf(nested).Len() != firstLength {
			return false
		}
	}
	return true
}
