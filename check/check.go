package check

import (
	"reflect"
)

// TODO check:
//  - Strict() & Loose() instead of Flag()
//  - tests
//  - more efficient (pointers etc.)

type Check[T any] interface {
	Get() T
	IsPresent() bool
	IsEmpty() bool
	IfPresent(method func(value T))
	IfEmpty(method func())
	IfPresentOrElse(method func(value T), altMethod func())
	OrElseGet(value T) T

	Map(method func(value T) T) check[T]
	MapToInt(method func(value T) int) check[int]
	MapToString(method func(value T) string) check[string]
	MapToBool(method func(value T) bool) check[bool]
	OrElse(value T) check[T]

	Peek(method func(value T)) check[T]
	Flag(flag ...flag) check[T]
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
