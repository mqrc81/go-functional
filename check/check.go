package check

import (
	"reflect"
)

// TODO check:
//  - tests
//  - more efficient (pointers etc.)

type Check[T any] interface {
	Map(mapFunc func(value T) T) check[T]
	MapToInt(mapFunc func(value T) int) check[int]
	MapToString(mapFunc func(value T) string) check[string]
	MapToBool(mapFunc func(value T) bool) check[bool]
	OrElse(value T) check[T]
	Peek(peekFunc func(value T)) check[T]
	Strict() check[T]
	Loose() check[T]

	Get() T
	IsPresent() bool
	IsEmpty() bool
	IfPresent(ifPresentFunc func(value T))
	IfEmpty(ifEmptyFunc func())
	IfPresentOrElse(ifPresentFunc func(value T), ifEmptyFunc func())
	OrElseGet(value T) T
}

type check[T any] struct {
	value  T
	valid  bool
	strict bool
}

func isValid(value any, strict bool) bool {
	rv := reflect.ValueOf(value)
	if strict {
		return isValidStrict(rv)
	}
	return isValidDefault(rv)
}

func isValidDefault(value reflect.Value) bool {
	return value.IsValid() && !value.IsZero()
}

func isValidStrict(value reflect.Value) bool {
	return isValidDefault(value) && !hasEmptyLength(value)
}

func hasEmptyLength(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan:
		return value.Len() == 0
	default:
		return false
	}
}
