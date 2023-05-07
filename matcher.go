package query

// Matcher is a function that takes an item and a key and returns a boolean. Return true if nil to match everything.
type Matcher[K comparable, V any] func(item V, key K) bool

// DoesMatch returns the result of the matcher.
func (m Matcher[K, V]) DoesMatch(item V, key K) bool {
	if m.isNil() {
		return true
	}

	return m(item, key)
}

// isNil returns true if the matcher is nil.
func (m Matcher[K, V]) isNil() bool {
	return m == nil
}

// Match returns a Matcher that matches items that match the given function.
func Match[K comparable, V any](m func(item V, key K) bool) Matcher[K, V] {
	return Matcher[K, V](m)
}
