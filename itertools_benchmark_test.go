package itertools

import (
	"math/rand"
	"reflect"
	"testing"
)

const repeatTimes = 1000
const stringLength = 100000

func BenchmarkIter(b *testing.B) {
	arr := rand.Perm(repeatTimes)
	for n := 0; n < b.N; n++ {
		counter := 0
		ch := Iter(arr)
		for range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Iter result not long enough")
			b.Fail()
		}
	}
}

func BenchmarkRepeatEdge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testString := generateRandomString(stringLength)
		ch := Repeat(testString, repeatTimes)
		counter := 0
		for value := range ch {
			if reflect.ValueOf(value).Len() == stringLength {
				counter++
			}
		}
		if counter != repeatTimes {
			b.Log("Repeat result not long enough")
			b.Fail()
		}
	}
}

func BenchmarkNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := rand.Perm(repeatTimes)
		counter := 0
		ch := Iter(arr)
		for counter < repeatTimes {
			counter++
			Next(ch)
		}
		if counter != repeatTimes {
			b.Log("Next result not long enough")
			b.Fail()
		}
	}
}

func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := 0
		first := rand.Perm(repeatTimes)
		second := rand.Perm(repeatTimes)
		third := rand.Perm(repeatTimes)
		ch := Zip(first, second, third)
		for range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Next result not long enough")
			b.Fail()
		}
	}
}

func BenchmarkAccumulate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "", 0)
		for range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}

	}
}
func BenchmarkAccumulateWithStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "", 100)
		for range ch {
			counter++
		}
		if counter-1 != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}
	}
}

func BenchmarkAccumulateMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "multiply", 0)
		for range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}
	}
}

func BenchmarkAccumulateMultiplyWithStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "multiply", 100)
		for range ch {
			counter++
		}
		if counter-1 != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}
	}
}

func BenchmarkTeeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := 0
		param := generateRandomString(stringLength)
		ch := Tee(param, 4)
		for range ch {
			counter++
		}
		if counter != stringLength/4 {
			b.Log("Tee results not long enough")
			b.Fail()
		}
	}
}
func BenchmarkTeeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(repeatTimes)
		counter := 0
		ch := Tee(arr, 4)
		for range ch {
			counter++
		}
		if counter != repeatTimes/4 {
			b.Log("Tee results not long enough")
			b.Fail()
		}
	}
}
