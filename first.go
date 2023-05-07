package query

// First returns the first item in the collection that matches the given matcher.
func First[K comparable, V any](col interface{}, matcher Matcher[K, V]) (K, V, error) {
	return Over[K, V](col).Filter(matcher).First()
}

// First returns the first item in the collection that matches the given matcher. or returns ErrNoMatches.
func (w *Walkable[K, V]) First() (K, V, error) {
	var (
		key   K
		value V
		found = false
	)

	_, err := w.Walk(func(item V, k K) error {
		//If we've already found a match, stop walking.
		key = k
		value = item
		found = true
		return ErrStopWalking
	})

	if err != nil {
		return key, value, err
	}

	if found {
		return key, value, nil
	}

	return key, value, ErrNoMatches
}
