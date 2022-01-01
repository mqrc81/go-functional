package stream

// TODO stream:
//  - everything
//  - Ordered flag; similar implementation as check.flag

type Stream[T any] interface {
	Collect(Collector[T]) []T
}

type stream[T any] struct {
	elements []T
}
