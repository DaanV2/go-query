package query

// Filter returns a new Walkable instance that filters out the items using the given matcher.
func Filter[K comparable, V any](coll interface{}, match Matcher[K, V]) *Walkable[K, V] {
	return Over[K, V](coll).Filter(match)
}

// Filter returns a new Walkable instance that filters out the items using the given matcher.
func (w *Walkable[K, V]) Filter(m Matcher[K, V]) *Walkable[K, V] {
	first := w.matcher
	w.matcher = first.And(m)
	return w
}
