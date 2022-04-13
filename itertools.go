package itertools

import (
	"math"
	"reflect"
	"strings"
)

type Iterator chan interface{}
type Predicate func(interface{}) bool
type Number interface {
	int | int8 | int32 | int64 | float32 | float64
}

// Iter returns an Iterator for the iterables parameter
func Iter[T any](iterables []T) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		for _, element := range iterables {
			ch <- element
		}
	}()
	return
}

// Next goes to the next item within an Iterator
func Next(ch Iterator) any {
	return <-ch
}

// Repeat returns an Iterator which contains element parameter, size parameter amount of times
func Repeat(element any, size int) Iterator {
	s := make([]any, size)
	for i := range s {
		s[i] = element
	}
	return Iter(s)
}

// Zip iterates over multiple data objects in sync
func Zip[T any](iterables ...[]T) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		if ok := ensureSameLength(iterables); !ok {
			ch <- "all parameters must be of the same length"
			return
		}
		var toSend []T
		for index := range iterables[0] {
			toSend = nil
			for _, iterable := range iterables {
				toSend = append(toSend, iterable[index])
			}
			ch <- toSend
		}
	}()
	return
}

// Chain allows for multiple arrays of the same type to be iterated over
func Chain[T any](iterables ...[]T) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		for _, iterable := range iterables {
			for index := range iterable {
				ch <- iterable[index]
			}
		}
	}()
	return
}

// Count counts up from a certain number in an increment
func Count[T Number](start, step T) (ch Iterator) {
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

// Accumulate returns an Iterator that sends/receives accumulated sums, or accumulated results of other functions
func Accumulate(iterable []int, operator string, start int) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		if start != 0 {
			ch <- start
		}
		toSend := iterable[0]
		ch <- toSend + start
		for _, element := range iterable[1:] {
			switch operator {
			case "add", "":
				toSend = toSend + element
			case "multiply":
				toSend = toSend * element
			case "power":
				toSend = int(math.Pow(float64(toSend), float64(element)))
			default:
				ch <- "not valid operator"
				return
			}
			ch <- toSend + start
		}
	}()
	return
}

// Tee returns the next n number of items next to each other
func Tee[T []int | string](iterable T, n int) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		switch reflect.TypeOf(iterable).Kind() {
		case reflect.String:
			element := reflect.ValueOf(iterable).String()
			for len(element) != 0 {
				if len(element) < n {
					ch <- element
					return
				}
				ch <- element[0:n]
				element = element[n:]
			}
		case reflect.Array, reflect.Slice:
			element := reflect.ValueOf(iterable)
			for element.Len() != 0 {
				if element.Len() < n {
					ch <- element
					return
				}
				toSend := element.Slice(0, n)
				element = element.Slice(n, element.Len())
				ch <- toSend
			}
		}
	}()
	return
}

// Pairwise sends/receives two characters next in a string via an Iterator
func Pairwise(iterable string) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		innerCh := Tee(iterable, 2)
		for element := range innerCh {
			ch <- element
		}
	}()
	return
}

// Dropwhile drops element from the iterable as long as the predicate is true - afterwards, returns every element using an Iterator
func Dropwhile[T Number](predicate Predicate, iterable []T) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		innerCh := Iter(iterable)
		for element := range innerCh {
			if !predicate(element) {
				ch <- element
				break
			}
		}
		for element := range innerCh {
			ch <- element
		}
	}()
	return
}

func Filterfalse[T Number](predicate Predicate, iterable []T) (ch Iterator) {
	ch = make(Iterator)
	go func() {
		defer close(ch)
		innerCh := Iter(iterable)
		for element := range innerCh {
			if !predicate(element) {
				ch <- element
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
