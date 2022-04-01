package itertools

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

func Next(ch Iterator) any {
	next := <-ch
	return next
}

func Repeat(value any, size int) Iterator {
	s := make([]any, size)
	for i := range s {
		s[i] = value
	}
	return Iter(s)
}

//func Zip[T any](iterables ...[]T) Iterator {
//	ch := make(Iterator)
//	go func() {
//		defer close(ch)
//		if ok := ensureSameLength(iterables); ok != true {
//			ch <- Result{Returned: false, Error: error.Error("all iterables are not that same length")}
//			break
//		}
//		var toSend []any
//		for index, _ := range iterables[0] {
//			toSend = nil
//			for _, iterable := range iterables {
//				toSend = append(toSend, iterable[index])
//			}
//			ch <- toSend
//		}
//	}()
//	return ch
//}
//
//func ensureSameLength[T any](nestedList [][]T) bool {
//	ch := Iter(nestedList)
//	first := Next(ch)
//	firstLength := reflect.ValueOf(first).Len()
//	fmt.Println(firstLength)
//	//firstLenght := len(first)
//	for nested := range ch {
//		if reflect.ValueOf(nested).Len() != firstLength {
//			return false
//		}
//	}
//}
