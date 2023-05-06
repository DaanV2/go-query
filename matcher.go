package query

// Matcher is a function that takes an item and a key and returns a boolean.
type Matcher[K comparable, V any] func(item V, key K) bool

func (m *Matcher[K, V]) DoesMatch(item V, key K) bool {
	if m == nil {
		return true
	}
	return (*m)(item, key)
}
