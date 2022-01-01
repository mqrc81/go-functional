package filter

import (
	"constraints"
)

func IsEven[T constraints.Integer](element T) bool {
	return element%2 == 0
}

func IsOdd[T constraints.Integer](element T) bool {
	return element%2 == 1
}

func IsDivisibleBy[T constraints.Integer](num T) func(element T) bool {
	return func(element T) bool {
		return element%num == 0
	}
}
