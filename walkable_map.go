package query

var _ IWalkable[int, int] = &WalkableArray[int]{}

// WalkableArray is a walkable array.
type WalkableMap[K comparable, V any] struct {
	items map[K]V
}

// NewWalkableMap creates a new walkable map.
func NewWalkableMap[K comparable, V any](items map[K]V) *WalkableMap[K, V] {
	return &WalkableMap[K, V]{
		items: items,
	}
}

// Walk walks over the items in the collection.
func (w *WalkableMap[K, V]) Walk(call func(item V, key K) error) (IWalkable[K, V], error) {
	for key, item := range w.items {
		if err := call(item, key); err != nil {
			return w, err
		}
	}
	return w, nil
}
