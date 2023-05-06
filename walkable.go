package query

type IWalkable[K any, V any] interface {
	// Walk walks over the items in the collection.
	Walk(call func(item V, key K) error) (IWalkable[K, V], error)
}

type Walkable[K comparable, V any] struct {
	source  IWalkable[K, V]
	matcher Matcher[K, V]
}

func (w *Walkable[K, V]) Walk(call func(item V, key K) error) (walker IWalkable[K, V], err error) {
	walker = w
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	_, err = w.source.Walk(func(item V, key K) error {
		if w.matcher.DoesMatch(item, key) {
			return call(item, key)
		}
		return nil
	})

	if err == ErrStopWalking {
		err = nil
	}

	return
}

func (w *Walkable[K, V]) ToArray() ([]V, error) {
	items := make([]V, 0)

	_, err := w.Walk(func(item V, key K) error {
		items = append(items, item)
		return nil
	})

	return items, err
}

func (w *Walkable[K, V]) ToMap() (map[K]V, error) {
	items := make(map[K]V)

	_, err := w.Walk(func(item V, key K) error {
		items[key] = item
		return nil
	})

	return items, err
}
