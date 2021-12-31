package check

func (c *check[T]) Get() T {
	return c.value
}

func (c *check[T]) IsPresent() bool {
	return c.valid
}

func (c *check[T]) IsEmpty() bool {
	return !c.IsPresent()
}

func (c *check[T]) IfPresent(method func(value T)) {
	if c.IsPresent() {
		method(c.value)
	}
}

func (c *check[T]) IfEmpty(method func()) {
	if c.IsEmpty() {
		method()
	}
}

func (c *check[T]) IfPresentOrElse(method func(value T), altMethod func()) {
	if c.IsPresent() {
		method(c.value)
	} else {
		altMethod()
	}
}

func (c *check[T]) OrElse(altValue T) *check[T] {
	if c.IsEmpty() {
		return &check[T]{
			altValue,
			isValid(altValue, c.strict),
			c.strict,
		}
	}
	return c
}

func (c *check[T]) OrElseGet(altValue T) T {
	if c.IsEmpty() {
		return altValue
	}
	return c.value
}
