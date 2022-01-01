package stream

func (s *stream[T]) Filter(filterFunc func(element T) bool) *stream[T] {
	s.applyFilter(filterFunc)
	return s
}

func (s *stream[T]) applyFilter(filterFunc func(element T) bool) {
	var toRemove []int
	for i, el := range s.elements {
		if !filterFunc(el) {
			toRemove = append(toRemove, i)
		}
	}
	s.elements = removeFromSlice(s.elements, toRemove)
}

func removeFromSlice[T any](slice []T, remove []int) []T {
	lenSlice, lenRemove := len(slice), len(remove)
	for i := len(remove) - 1; i >= 0; i-- {
		slice[remove[i]] = slice[lenSlice-(lenRemove-i)]
	}
	return slice[:len(slice)-len(remove)]
}
