package check

func Of[T any](value T, options ...flag) *check[T] {
	strict, valid := fromFlags(options)

	return &check[T]{
		value:  value,
		valid:  valid,
		strict: strict,
	}
}

func Empty[T any](options ...flag) *check[T] {
	strict, _ := fromFlags(options)

	return &check[T]{
		strict: strict,
	}
}

type flag int

const (
	_ flag = iota

	// Valid states that the value passed into check.Of() is not empty,
	// thus allowing to skip the initial empty-check.
	// Flag is ignored for check.Empty().
	Valid

	// Strict enables strict-mode. With strict-mode, additionally,
	// a slice, map or chan of length 0 is treated as empty.
	// Strict-mode is disabled by default.
	Strict
)

func fromFlags(options []flag) (valid bool, strict bool) {
	for _, option := range options {
		switch option {
		case Valid:
			valid = true
		case Strict:
			strict = true
		}
	}
	return valid, strict
}
