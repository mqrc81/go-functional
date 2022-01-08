package stream

import (
	"go-functional/check"
)

func (s stream[T]) Collect() []T {
	var result []T
	for _, element := range s.elements {
		if s.terminate(&element) {
			result = append(result, element)
		}
	}
	return result
}

func (s stream[T]) Count() int {
	var count int
	for _, element := range s.elements {
		if s.terminate(&element) {
			count++
		}
	}
	return count
}

func (s stream[T]) AnyMatch(matchFunc func(element T) bool) bool {
	for _, element := range s.elements {
		if s.terminate(&element) {
			if matchFunc(element) {
				return true
			}
		}
	}
	return false
}

func (s stream[T]) AllMatch(matchFunc func(element T) bool) bool {
	for _, element := range s.elements {
		if s.terminate(&element) {
			if !matchFunc(element) {
				return false
			}
		}
	}
	return true
}

func (s stream[T]) NoneMatch(matchFunc func(element T) bool) bool {
	for _, element := range s.elements {
		if s.terminate(&element) {
			if matchFunc(element) {
				return false
			}
		}
	}
	return true
}

func (s stream[T]) ForEach(forEachFunc func(element T)) {
	for _, element := range s.elements {
		if s.terminate(&element) {
			forEachFunc(element)
		}
	}
}

func (s stream[T]) Find(matchFunc func(element T) bool) check.Check[T] {
	for _, element := range s.elements {
		if s.terminate(&element) {
			if matchFunc(element) {
				return check.Of[T](element, check.Valid)
			}
		}
	}
	return check.Empty[T]()
}

func (s stream[T]) Fold(initialValue T, foldFunc func(value *T, element T)) T {
	var result = initialValue
	for _, element := range s.elements {
		if s.terminate(&element) {
			foldFunc(&result, element)
		}
	}
	return result
}

func (s stream[T]) FoldToInt(initialValue int, foldFunc func(value *int, element T)) int {
	var result = initialValue
	for _, element := range s.elements {
		if s.terminate(&element) {
			foldFunc(&result, element)
		}
	}
	return result
}

func (s stream[T]) FoldToString(initialValue string, foldFunc func(value *string, element T)) string {
	var result = initialValue
	for _, element := range s.elements {
		if s.terminate(&element) {
			foldFunc(&result, element)
		}
	}
	return result
}

func (s *stream[T]) terminate(element *T) bool {
	if len(s.operations) == 0 {
		return true
	}
	return s.operations[0].apply(element, &s.operations, 0, s.parallel)
}
