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
