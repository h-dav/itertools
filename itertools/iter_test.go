package itertools

import (
	"fmt"
)

func ExampleIterInt() {
	arr := []int{1, 2, 3, 4}
	ch := Iter(arr)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: 1234
}

func ExampleNext() {
	arr := []int{1, 2, 3, 4}
	ch := Iter(arr)
	value := Next(ch)
	fmt.Printf("%v", value)
	value = Next(ch)
	fmt.Printf("%v", value)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: 1234
}

func ExampleRepeat() {
	ch := Repeat("hello", 5)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: hellohellohellohellohello
}

func ExampleZip() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
	ch := Zip(first, second, third)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: [1 4 7][2 5 8][3 6 9]
}
func ExampleZipFailure() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9, 11}
	ch := Zip(first, second, third)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: all parameters must be of the same length
}
