package query

// Count returns the number of items in the collection.
func (w *Walkable[K, V]) Count() (int, error) {
	count := 0
	_, err := w.Walk(func(item V, key K) error {
		count++
		return nil
	})

	return count, err
}
