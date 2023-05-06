package query

// Any returns a matcher that matches if any matcher matches.
func Any[K comparable, V any](matcher ...Matcher[K, V]) Matcher[K, V] {
	return func(item V, key K) bool {
		for _, m := range matcher {
			if m(item, key) {
				return true
			}
		}
		return false
	}
}

// Any returns a matcher that matches if any matcher matches.
func (m *Matcher[K, V]) Any(matcher ...Matcher[K, V]) Matcher[K, V] {
	return m.Or(Any(matcher...))
}

// Any returns a matcher that matches if any matcher matches.
func (w *Walkable[K, V]) Any(matcher ...Matcher[K, V]) *Walkable[K, V] {
	first := w.matcher
	w.matcher = first.Any(matcher...)
	return w
}

// HasAny returns a boolean indicating whether any of the matchers match.
func HasAny[K comparable, V any](col interface{}, any ...Matcher[K, V]) (bool, error) {
	return Over[K, V](col).Any(any...).HasAny()
}

// HasAny returns a boolean indicating whether any of the matchers match.
func (w *Walkable[K, V]) HasAny() (bool, error) {
	found := true
	_, _, err := w.First()
	if err == ErrNoMatches {
		err = nil
	}

	return found, err
}
