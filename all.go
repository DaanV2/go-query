package query

// All returns a matcher that matches if all matchers match.
func All[K comparable, V any](matcher ...Matcher[K, V]) Matcher[K, V] {
	return func(item V, key K) bool {
		for _, m := range matcher {
			if !m(item, key) {
				return false
			}
		}

		return true
	}
}

// All returns a matcher that matches if all matchers match.
func (m *Matcher[K, V]) All(matcher ...Matcher[K, V]) Matcher[K, V] {
	return m.And(All(matcher...))
}

// All returns a matcher that matches if all matchers match.
func (w *Walkable[K, V]) All(matcher ...Matcher[K, V]) *Walkable[K, V] {
	first := w.matcher
	w.matcher = first.All(matcher...)
	return w
}

// HasAll returns a boolean indicating whether a single item matchers all of the match.
func HasAll[K comparable, V any](col interface{}, any ...Matcher[K, V]) (bool, error) {
	return Over[K, V](col).All(any...).HasAny()
}
