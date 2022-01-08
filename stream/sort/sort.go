package sort

import (
	"constraints"
)

// These helper functions can be used for: stream.Sort()

func Ascending[T constraints.Ordered](element1 T, element2 T) bool {
	return element1 > element2
}

func Descending[T constraints.Ordered](element1 T, element2 T) bool {
	return element1 < element2
}

func LenAscending[T lengthable](element1 T, element2 T) bool {
	return len(element1) > len(element2)
}

func LenDescending[T lengthable](element1 T, element2 T) bool {
	return len(element1) < len(element2)
}

type lengthable interface {
	~string |
		~[]any | ~[]int | ~[]string | ~[]bool | ~[]float32 | ~[]int64 | ~[]int32 | ~[]int16 | ~[]int8 | ~[]byte |
		~map[any]any | ~map[string]any | ~map[string]int | ~map[string]bool | ~map[int]any | ~map[int]string | ~map[int]bool |
		~chan any | ~chan string | ~chan int | ~chan bool | ~chan float32
}
