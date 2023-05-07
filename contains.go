package query

// Contains checks if any item in the collection matches the given matcher.
func Contains[K comparable, V any](coll interface{}, match Matcher[K, V]) (bool, error) {
	source := Over[K, V](coll)
	found := false

	_, err := source.Walk(func(item V, key K) error {
		if match.DoesMatch(item, key) {
			found = true
			return ErrStopWalking
		}

		return nil
	})

	return found, err
}
