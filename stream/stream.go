package stream

// TODO stream:
//  - everything

type Stream interface {
}

type stream[T any] struct {
	values   []T
	sequence []action
}

type action struct {
	// TODO
}
