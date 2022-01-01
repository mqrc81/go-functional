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
