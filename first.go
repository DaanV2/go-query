package query

// First returns the first item that matches the given matcher.
func First[K comparable, V any](col interface{}, matcher Matcher[K, V]) (K, V, error) {
	return Over[K, V](col).First(matcher)
}

// First returns the first item that matches the given matcher. or returns ErrNoMatches.
func (w *Walkable[K, V]) First(match Matcher[K, V]) (K, V, error) {
	var (
		key   K
		value V
		found = false
	)

	_, err := w.Walk(func(item V, k K) error {
		if match.DoesMatch(item, k) {
			key = k
			value = item
			found = true
			return ErrStopWalking
		}

		return nil
	})

	if err != nil {
		return key, value, err
	}

	if found {
		return key, value, nil
	}

	return key, value, ErrNoMatches
}
