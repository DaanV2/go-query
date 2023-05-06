package query

var _ IWalkable[int, int] = &WalkableArray[int]{}

// WalkableArray is a walkable array.
type WalkableArray[V any] struct {
	arr []V
}

// NewWalkableArray creates a new walkable array.
func NewWalkableArray[V any](arr []V) *WalkableArray[V] {
	return &WalkableArray[V]{
		arr: arr,
	}
}

// Walk walks over the items in the collection.
func (w *WalkableArray[V]) Walk(call func(item V, key int) error) (IWalkable[int, V], error) {
	for key, item := range w.arr {
		if err := call(item, key); err != nil {
			return IWalkable[int, V](w), err
		}
	}
	return IWalkable[int, V](w), nil
}
