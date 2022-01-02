package check

func (c check[T]) Peek(method func(value T)) check[T] {
	method(c.value)
	return c
}
