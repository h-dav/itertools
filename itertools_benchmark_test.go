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
