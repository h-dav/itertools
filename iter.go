package operation

import (
	"reflect"
	"strings"
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
		if ok := ensureSameLength(iterables); !ok {
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

// Chain allows for multiple arrays of the same type to be iterated over
func Chain[T any](iterables ...[]T) Iterator {
	ch := make(Iterator)
	go func() {
		defer close(ch)
		for _, iterable := range iterables {
			for index, _ := range iterable {
				ch <- iterable[index]
			}
		}
	}()
	return ch
}

//func Combinations(iterable string, times uint) Iterator {
//	// TODO
//	ch := make(Iterator)
//	go func() {
//		defer close(ch)
//		prepared := strings.SplitAfter(iterable, "")
//		for index, value := range prepared {
//			iterator := Iter(prepared[index:])
//			for nextVal := range iterator {
//				ch <- value + fmt.Sprintf("%v", nextVal)
//			}
//		}
//	}()
//	return ch
//}

// Count counts up from a certain number in an increment
//func Count[T, S float32 | float64 | int](start T, step S) (ch Iterator) {
func Count[T float32 | float64 | int](start, step T) (ch Iterator) {
	// consider changing step to uint
	ch = make(Iterator)
	go func() {
		defer close(ch)
		for {
			ch <- start
			start = start + step
		}
	}()
	return
}

// Cycle goes over a string seemingly forever
func Cycle(iterable string) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		for {
			letters := strings.SplitAfter(iterable, "")
			for _, letter := range letters {
				ch <- letter
			}
		}
	}()
	return
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
