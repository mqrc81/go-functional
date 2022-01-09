package stream

func (s stream[T]) Filter(filterFunc func(element T) bool) stream[T] {
	s.operations = append(s.operations, filterOperation[T]{
		filterFunc: filterFunc,
	})
	return s
}

func (s stream[T]) Limit(maxSize int) stream[T] {
	s.operations = append(s.operations, limitOperation[T]{
		maxSize:    maxSize,
		spotsTaken: new(int),
	})
	return s
}

func (s stream[T]) Reverse() stream[T] {
	if s.parallel {
		return s
	}
	var reversed []T
	for _, element := range s.elements {
		if s.terminate(&element) {
			prepend(&reversed, element)
		}
	}
	s.elements = reversed
	s.operations = nil
	return s
}

func (s stream[T]) Sort(sort1Above2Func func(element1 T, element2 T) bool) stream[T] {
	var sorted []T
	for _, element := range s.elements {
		if s.terminate(&element) {
			insertSorted(&sorted, &element, sort1Above2Func)
		}
	}
	s.elements = sorted
	s.operations = nil
	return s
}

func (s stream[T]) Map(mapFunc func(element T) T) stream[T] {
	s.operations = append(s.operations, mapOperation[T]{
		mapFunc: mapFunc,
	})
	return s
}

func (s stream[T]) MapToInt(mapFunc func(element T) int) stream[int] {
	var mapped []int
	for _, element := range s.elements {
		if s.terminate(&element) {
			mapped = append(mapped, mapFunc(element))
		}
	}
	return stream[int]{
		elements: mapped,
		parallel: s.parallel,
	}
}

func (s stream[T]) MapToString(mapFunc func(element T) string) stream[string] {
	var mapped []string
	for _, element := range s.elements {
		if s.terminate(&element) {
			mapped = append(mapped, mapFunc(element))
		}
	}
	return stream[string]{
		elements: mapped,
		parallel: s.parallel,
	}
}

func (s stream[T]) TakeWhile(takeWhileFunc func(element T) bool) stream[T] {
	s.operations = append(s.operations, takeWhileOperation[T]{
		takeWhileFunc: takeWhileFunc,
		stopTaking:    new(bool),
	})
	return s
}

func (s stream[T]) DropWhile(dropWhileFunc func(element T) bool) stream[T] {
	s.operations = append(s.operations, dropWhileOperation[T]{
		dropWhileFunc: dropWhileFunc,
		stopDropping:  new(bool),
	})
	return s
}

func (s stream[T]) Peek(peekFunc func(element T)) stream[T] {
	s.operations = append(s.operations, peekOperation[T]{
		peekFunc: peekFunc,
	})
	return s
}

func (s stream[T]) Parallel() stream[T] {
	s.parallel = true
	return s
}

func (s stream[T]) Ordered() stream[T] {
	s.parallel = false
	return s
}

func prepend[T any](slice *[]T, element T) {
	*slice = append(*slice, element)
	copy((*slice)[1:], *slice)
	(*slice)[0] = element
}

func insertSorted[T any](sorted *[]T, element *T, sortFunc func(el1 T, el2 T) bool) {
	*sorted = append(*sorted, *element)
	for i := 0; i < len(*sorted)-1; i++ {
		if !sortFunc(*element, (*sorted)[i]) {
			copy((*sorted)[i+1:], (*sorted)[i:])
			(*sorted)[i] = *element
			break
		}
	}
}
