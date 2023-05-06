package query

var _ IWalkable[int, int] = &WalkableArray[int]{}

type WalkableMap[K comparable, V any] struct {
	items map[K]V
}

func NewWalkableMap[K comparable, V any](items map[K]V) *WalkableMap[K, V] {
	return &WalkableMap[K, V]{
		items: items,
	}
}

func (w *WalkableMap[K, V]) Walk(call func(item V, key K) error) (IWalkable[K, V], error) {
	for key, item := range w.items {
		if err := call(item, key); err != nil {
			return IWalkable[K, V](w), err
		}
	}
	return IWalkable[K, V](w), nil
}
