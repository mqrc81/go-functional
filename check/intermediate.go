package check

// Go doesn't allow parameterized methods, thus we can't map from T to any type U
// and I refuse to pass check as a parameter because it's ugly/inconsistent)
// (see: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods)

func (c check[T]) Map(method func(value T) T) check[T] {
	return mapOrEmpty(&c, method)
}

func (c check[T]) MapToInt(method func(value T) int) check[int] {
	return mapOrEmpty(&c, method)
}

func (c check[T]) MapToString(method func(value T) string) check[string] {
	return mapOrEmpty(&c, method)
}

func (c check[T]) MapToBool(method func(value T) bool) check[bool] {
	return mapOrEmpty(&c, method)
}

func mapOrEmpty[T, U any](c *check[T], method func(value T) U) check[U] {
	var newC check[U]
	newC.valid = c.IsPresent()
	newC.strict = c.strict
	if newC.valid {
		newC.value = method(c.value)
	}
	return newC
}

func (c check[T]) OrElse(altValue T) check[T] {
	if c.IsEmpty() {
		return check[T]{
			altValue,
			isValid(altValue, c.strict),
			c.strict,
		}
	}
	return c
}

func (c check[T]) Peek(method func(value T)) check[T] {
	method(c.value)
	return c
}

func (c check[T]) Strict() check[T] {
	c.strict = true
	return c
}

func (c check[T]) Loose() check[T] {
	c.strict = false
	return c
}
