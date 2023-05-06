package query

// Matcher is a function that takes an item and a key and returns a boolean. Return true if nil to match everything.
type Matcher[K comparable, V any] func(item V, key K) bool

// DoesMatch returns the result of the matcher.
func (m Matcher[K, V]) DoesMatch(item V, key K) bool {
	if m == nil {
		return true
	}
	return m(item, key)
}

// Match returns a Matcher that matches items that match the given function.
func Match[K comparable, V any](m func(item V, key K) bool) Matcher[K, V] {
	return Matcher[K, V](m)
}
