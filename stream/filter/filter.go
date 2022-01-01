package filter

import (
	"constraints"
)

// These helper functions can be used for: stream.Filter(), stream.AnyMatch, stream.AllMatch()

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
