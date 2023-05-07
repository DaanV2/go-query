package query

// Not returns a matcher that matches if A does not match.
func Not[K comparable, V any](A Matcher[K, V]) Matcher[K, V] {
	return func(item V, key K) bool {
		return !A.DoesMatch(item, key)
	}
}

// Not returns a matcher that matches if A does not match.
func (m *Matcher[K, V]) Not() Matcher[K, V] {
	return Not(*m)
}

// Not returns a matcher that matches if A does not match.
func (w *Walkable[K, V]) Not() *Walkable[K, V] {
	first := w.matcher
	w.matcher = first.Not()
	return w
}
