package query

type IWalkable[K any, V any] interface {
	// Walk walks over the items in the collection.
	Walk(call func(item V, key K) error) (IWalkable[K, V], error)
}

type Walkable[K any, V any] struct {
	source IWalkable[K, V]
}

func (w *Walkable[K, V]) Walk(call func(item V, key K) error) (IWalkable[K, V], error) {
	return w.source.Walk(call)
}
