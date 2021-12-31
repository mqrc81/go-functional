package check

// Go doesn't allow functions with receivers to declare type parameters,
// so we can't map T to another generic type U
// (and I refuse to pass check as a parameter because it's ugly/inconsistent)
// For now, MapToInt & MapToString exist, but to unveil Check's full potential,
// hopefully Go will support this use-case with public release of Go 1.18

func (c *check[T]) Map(method func(value T) T) *check[T] {
	return mapOrEmpty(c, method)
}

func (c *check[T]) MapToInt(method func(value T) int) *check[int] {
	return mapOrEmpty(c, method)
}

func (c *check[T]) MapToString(method func(value T) string) *check[string] {
	return mapOrEmpty(c, method)
}

func (c *check[T]) MapToBool(method func(value T) bool) *check[bool] {
	return mapOrEmpty(c, method)
}

func mapOrEmpty[T, U any](c *check[T], method func(value T) U) *check[U] {
	var newC check[U]
	newC.valid = c.IsPresent()
	newC.strict = c.strict
	if newC.valid {
		newC.value = method(c.value)
	}
	return &newC
}
