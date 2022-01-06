package stream

func (s stream[T]) Collect() []T {
	var result []T
	for _, element := range s.elements {
		if s.operations[0].apply(&element, &s.operations, 0, s.ordered) {
			result = append(result, element)
		}
	}
	return result
}
