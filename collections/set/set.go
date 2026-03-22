package set

// Set is a generic set implementation backed by a map.
type Set[T comparable] map[T]struct{}

// New creates a new set and optionally fills it with values.
func New[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	set.Add(values...)
	return set
}

// Add inserts values into the set.
func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

// Delete removes values from the set.
func (s Set[T]) Delete(values ...T) {
	for _, value := range values {
		delete(s, value)
	}
}

// Exists reports whether the value exists in the set.
func (s Set[T]) Exists(value T) bool {
	_, ok := s[value]
	return ok
}

// Get returns the value and whether it exists in the set.
func (s Set[T]) Get(value T) (T, bool) {
	_, ok := s[value]
	return value, ok
}

// Len returns the number of values in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Clear removes all values from the set.
func (s Set[T]) Clear() {
	for value := range s {
		delete(s, value)
	}
}

// Values returns all set values as a slice.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for value := range s {
		values = append(values, value)
	}
	return values
}
