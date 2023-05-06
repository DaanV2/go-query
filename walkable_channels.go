package query

var _ IWalkable[int, int] = &WalkableChannel[int]{}

// IWalkableChannel is a walkable channel.
type WalkableChannel[V any] struct {
	ch chan V
}

// NewWalkableChan creates a new walkable channel.
func NewWalkableChan[V any](ch chan V) *WalkableChannel[V] {
	return &WalkableChannel[V]{
		ch: ch,
	}
}

// Walk walks over the items in the collection.
func (w *WalkableChannel[V]) Walk(call func(item V, key int) error) (IWalkable[int, V], error) {
	i := 0
	for item := range w.ch {
		if err := call(item, i); err != nil {
			return w, err
		}
		i++
	}
	return w, nil
}
