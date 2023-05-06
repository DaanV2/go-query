package query

type WalkableChannel[V any] struct {
	ch chan V
}

var _ IWalkable[int, int] = &WalkableChannel[int]{}

func NewWalkableChan[V any](ch chan V) *WalkableChannel[V] {
	return &WalkableChannel[V]{
		ch: ch,
	}
}

// Walk implements IWalkable
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
