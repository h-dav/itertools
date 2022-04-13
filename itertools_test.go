package itertools

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleIterInt() {
	arr := []int{1, 2, 3, 4}
	ch := Iter(arr)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: 1:2:3:4:
}

func TestIterLower(t *testing.T) {
	expected, counter := 1, 0
	arr := []int{1}
	ch := Iter(arr)
	for range ch {
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
	element := Next(ch)
	fmt.Printf("%v", element)
	element = Next(ch)
	fmt.Printf("%v", element)
	for element := range ch {
		fmt.Printf("%v", element)
	}
	// Output: 1234
}

func ExampleRepeat() {
	ch := Repeat("example_string", 5)
	for element := range ch {
		fmt.Printf("%v", element)
	}
	// Output: example_stringexample_stringexample_stringexample_stringexample_string
}

func ExampleZip() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
	ch := Zip(first, second, third)
	for element := range ch {
		fmt.Printf("%v", element)
	}
	// Output: [1 4 7][2 5 8][3 6 9]
}

func ExampleZipFailure() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9, 11}
	ch := Zip(first, second, third)
	for element := range ch {
		fmt.Printf("%v", element)
	}
	// Output: all parameters must be of the same length
}

func ExampleChain() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
	ch := Chain(first, second, third)
	for element := range ch {
		fmt.Printf("%v", element)
	}
	// Output: 123456789
}

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
		fmt.Printf("%v:", Next(ch))
	}
	// Output:1.5:2:2.5:3:
}

func ExampleCycle() {
	ch := Cycle("teststring")
	for i := 0; i < 14; i++ {
		fmt.Printf("%v", Next(ch))
	}
	// Output: teststringtest
}

func ExampleAccumulate() {
	arr := []int{1, 2, 3, 4, 5}
	ch := Accumulate(arr, "", 0)
	for element := range ch {
		fmt.Printf("%v", element)
	}
	// Output: 1361015
}

func ExampleAccumulateWithStart() {
	arr := []int{1, 2, 3, 4, 5}
	ch := Accumulate(arr, "", 100)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: 100:101:103:106:110:115:
}

func ExampleAccumulateMultiply() {
	arr := []int{1, 2, 3, 4, 5}
	ch := Accumulate(arr, "multiply", 0)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: 1:2:6:24:120:
}

func ExampleAccumulateMultiplyWithStart() {
	arr := []int{1, 2, 3, 4, 5}
	ch := Accumulate(arr, "multiply", 100)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: 100:101:102:106:124:220:
}

func ExampleTee() {
	param := "ABCDEFGHIJKLMNOPQ"
	ch := Tee(param, 2)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: AB:CD:EF:GH:IJ:KL:MN:OP:Q:
}

func ExampleTeeArray() {
	param := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := Tee(param, 2)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: [1 2]:[3 4]:[5 6]:[7 8]:[9]:
}

func ExamplePairwise() {
	param := "ABCDEFGHIJKLP"
	ch := Pairwise(param)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: AB:CD:EF:GH:IJ:KL:P:
}

func ExampleDropwhile() {
	param := []int{1, 4, 6, 4, 1}
	predicate := Predicate(func(i interface{}) bool {
		return i.(int) < 5
	})
	ch := Dropwhile(predicate, param)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: 6:4:1:
}

func ExampleFilterfalse() {
	param := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isOdd := Predicate(func(i interface{}) bool {
		result := i.(int) % 2
		if result == 1 {
			return true
		}
		return false
	})
	ch := Filterfalse(isOdd, param)
	for element := range ch {
		fmt.Printf("%v:", element)
	}
	// Output: 2:4:6:8:10:
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
