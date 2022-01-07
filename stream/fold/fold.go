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
func Max() (int, func(value *int, element int)) {
	return minInt, func(value *int, element int) {
		if element > *value {
			*value = element
		}
	}
}

// TODO generify
func Min() (int, func(value *int, element int)) {
	return maxInt, func(value *int, element int) {
		if element < *value {
			*value = element
		}
	}
}

func Sum[T Number]() (T, func(value *T, element T)) {
	return 0, func(value *T, element T) {
		*value += element
	}
}

func Product[T Number]() (T, func(value *T, element T)) {
	return 1, func(value *T, element T) {
		*value *= element
	}
}

type Number interface {
	constraints.Integer | constraints.Float
}

// TODO add delimiter
func Concat() (string, func(value *string, element string)) {
	return "", func(value *string, element string) {
		*value += element
	}
}
