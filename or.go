package query

// Or returns a matcher that matches if either A or B match.
func Or[K comparable, V any](A, B Matcher[K, V]) Matcher[K, V] {
	return func(item V, key K) bool {
		return A.DoesMatch(item, key) || B.DoesMatch(item, key)
	}
}

// Or returns a matcher that matches if either A or B match.
func (m *Matcher[K, V]) Or(or Matcher[K, V]) Matcher[K, V] {
	if m == nil {
		return or
	}

	return Or(*m, or)
}

// Or returns a matcher that matches if either A or B match.
func (w *Walkable[K, V]) Or(and Matcher[K, V]) *Walkable[K, V] {
	first := w.matcher
	w.matcher = Or(first, and)
	return w
}
