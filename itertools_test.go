package itertools

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleIterInt() {
	arr := []int{1, 2, 3, 4}
	ch := Iter(arr)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: 1234
}

func TestIterLower(t *testing.T) {
	expected, counter := 1, 0
	arr := []int{1}
	ch := Iter(arr)
	for _ = range ch {
		counter++
	}
	if counter != expected {
		t.Log("counter not correct number")
		t.Fail()
	}
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
	ch := Repeat("example_string", 5)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: example_stringexample_stringexample_stringexample_stringexample_string
}

func generateRandomString(stringLength int) (result string) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, stringLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	result = string(b)
	return
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

func ExampleChain() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
	ch := Chain(first, second, third)
	for value := range ch {
		fmt.Printf("%v", value)
	}
	// Output: 123456789
}

//func ExampleCombinations() {
//	ch := Combinations("abcd", 3)
//	for value := range ch {
//		fmt.Printf("%v", value)
//	}
//	// Output: abacad
//}

func ExampleCount() {
	ch := Count(1, 2)
	fmt.Printf("%v", Next(ch))
	fmt.Printf("%v", Next(ch))
	fmt.Printf("%v", Next(ch))
	// Output: 135
}

func ExampleCountDec() {
	ch := Count(1.5, 0.5)
	for i := 0; i < 4; i++ {
		fmt.Printf("%v,", Next(ch))
	}
	// Output:1.5,2,2.5,3,
}

func ExampleCycle() {
	ch := Cycle("teststring")
	for i := 0; i < 14; i++ {
		fmt.Printf("%v", Next(ch))
	}
	// Output: teststringtest
}