package option

type Option[T any] struct {
	value *T
}

func Of[T any](value T) *Option[T] {
	return &Option[T]{
		&value,
	}
}

func Empty[T any]() *Option[T] {
	return &Option[T]{
		nil,
	}
}

func (opt *Option[T]) IsPresent() bool {
	return !opt.isNil()
}

func (opt *Option[T]) IfPresent(method func(value T)) {
	if !opt.isNil() {
		method(*opt.value)
	}
}

// Go doesn't allow functions with receivers to declare type parameters, thus we can't map T to U
// (and I refuse to pass Option as a parameter because it's ugly/inconsistent)
func (opt *Option[T]) Map(method func(value T) T) *Option[T] {
	if opt.isNil() {
		return opt
	}
	newValue := method(*opt.value)
	return &Option[T]{
		&newValue,
	}
}

func (opt *Option[T]) Get() T {
	return *opt.value
}

func (opt *Option[T]) OrElse(alternativeValue T) T {
	if opt.isNil() {
		return alternativeValue
	}
	return *opt.value
}

func (opt *Option[T]) Or(alternativeValue T) *Option[T] {
	if opt.isNil() {
		return &Option[T]{
			&alternativeValue,
		}
	}
	return opt
}

func (opt *Option[T]) OrElseGet(alternativeOption *Option[T]) T {
	return opt.OrElse(*alternativeOption.value)
}

func (opt *Option[T]) isNil() bool {
	return opt.value == nil
}
