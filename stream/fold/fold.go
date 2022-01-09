package fold

import (
	"constraints"
)

// These helper functions can be used for: stream.Fold(), stream.FoldToInt(), stream.FoldToString()

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// TODO generify
func Max() (int, func(value int, element int) int) {
	return minInt, func(value int, element int) int {
		if element > value {
			value = element
		}
		return value
	}
}

// TODO generify
func Min() (int, func(value int, element int) int) {
	return maxInt, func(value int, element int) int {
		if element < value {
			value = element
		}
		return value
	}
}

func Sum[T Number]() (T, func(value T, element T) T) {
	return 0, func(value T, element T) T {
		return value + element
	}
}

func Product[T Number]() (T, func(value T, element T) T) {
	return 1, func(value T, element T) T {
		return value * element
	}
}

type Number interface {
	constraints.Integer | constraints.Float
}
