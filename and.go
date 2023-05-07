package query

// And returns a matcher that matches if both A and B match.
func And[K comparable, V any](A, B Matcher[K, V]) Matcher[K, V] {
	return func(item V, key K) bool {
		return A.DoesMatch(item, key) && B.DoesMatch(item, key)
	}
}

// And returns a matcher that matches if both A and B match.
func (m *Matcher[K, V]) And(and Matcher[K, V]) Matcher[K, V] {
	if m.isNil() {
		return and
	}

	return And(*m, and)
}

// And returns a matcher that matches if both A and B match.
func (w *Walkable[K, V]) And(and Matcher[K, V]) *Walkable[K, V] {
	first := w.matcher
	w.matcher = And(first, and)
	return w
}

