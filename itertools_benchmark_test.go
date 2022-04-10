package itertools

import (
	"math/rand"
	"reflect"
	"testing"
)

func BenchmarkIter(b *testing.B) {
	repeatTimes := 10000000
	arr := rand.Perm(repeatTimes)
	for n := 0; n < b.N; n++ {
		counter := 0
		ch := Iter(arr)
		for _ = range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Iter result not long enough")
			b.Fail()
		}
	}
}

func BenchmarkRepeatEdge(b *testing.B) {
	stringLength, repeatTimes := 100000, 100000000
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
	repeatTimes := 100000000
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
	repeatTimes := 100000000
	for i := 0; i < b.N; i++ {
		counter := 0
		first := rand.Perm(repeatTimes)
		second := rand.Perm(repeatTimes)
		third := rand.Perm(repeatTimes)
		ch := Zip(first, second, third)
		for _ = range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Next result not long enough")
			b.Fail()
		}
	}
}

func BenchmarkAccumulate(b *testing.B) {
	repeatTimes := 10000000
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "", 0)
		for _ = range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}

	}
}
func BenchmarkAccumulateWithStart(b *testing.B) {
	repeatTimes := 10000000
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "", 100)
		for _ = range ch {
			counter++
		}
		if counter-1 != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}
	}
}

func BenchmarkAccumulateMultiply(b *testing.B) {
	repeatTimes := 10000000
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "multiply", 0)
		for _ = range ch {
			counter++
		}
		if counter != repeatTimes {
			b.Log("Accumulate results not long enough")
			b.Fail()
		}
	}
}

func BenchmarkAccumulateMultiplyWithStart(b *testing.B) {
	repeatTimes := 10000000
	for i := 0; i < b.N; i++ {
		counter := 0
		arr := rand.Perm(repeatTimes)
		ch := Accumulate(arr, "multiply", 100)
		for _ = range ch {
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
		stringLength := 1000000
		counter := 0
		param := generateRandomString(stringLength)
		ch := tee(param, 4)
		for _ = range ch {
			counter++
		}
		if counter != stringLength/4 {
			b.Log("tee results not long enough")
			b.Fail()
		}
	}
}
func BenchmarkTeeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeatTimes := 1000000
		arr := rand.Perm(repeatTimes)
		counter := 0
		ch := tee(arr, 4)
		for _ = range ch {
			counter++
		}
		if counter != repeatTimes/4 {
			b.Log("tee results not long enough")
			b.Fail()
		}
	}
}
