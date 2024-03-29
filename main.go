package query

import (
	"errors"
	"fmt"
	"reflect"
)

// Over returns a new Walkable instance over the given collection. Might return an error if the collection is not supported.
func Over[K comparable, V any](coll interface{}) *Walkable[K, V] {
	source, err := generateSource[K, V](coll)
	if err != nil {
		source = NewWalkError[K, V](err)
	}
	return &Walkable[K, V]{
		source:  source,
		matcher: nil,
	}
}

// generateSource generates the source for the walkable from an interface.
func generateSource[K comparable, V any](coll interface{}) (IWalkable[K, V], error) {
	switch c := coll.(type) {
	case IWalkable[K, V]:
		return c, nil

	case []V:
		arr := IWalkable[int, V](NewWalkableArray(c))
		if r, ok := arr.(IWalkable[K, V]); ok {
			return r, nil
		}

		var k K
		keyName := reflect.TypeOf(k).Name()
		return nil, errors.New(fmt.Sprintf("cannot convert %T for array, need int", keyName))

	case chan V:
		chn := IWalkable[int, V](NewWalkableChan(c))
		if r, ok := chn.(IWalkable[K, V]); ok {
			return r, nil
		}

		var k K
		keyName := reflect.TypeOf(k).Name()
		return nil, errors.New(fmt.Sprintf("cannot convert %T for channel, need int", keyName))

	case map[K]V:
		return NewWalkableMap(c), nil

	default:
		return nil, ErrWrongWalkableType
	}
}

// OverArray returns a new Walkable instance over the given array.
func OverArray[V any](arr []V) *Walkable[int, V] {
	return &Walkable[int, V]{
		source: NewWalkableArray(arr),
	}
}

// OverChannel returns a new Walkable instance over the given channel.
func OverChannel[V any](chn chan V) *Walkable[int, V] {
	return &Walkable[int, V]{
		source: NewWalkableChan(chn),
	}
}

// OverMap returns a new Walkable instance over the given channel.
func OverMap[K comparable, V any](items map[K]V) *Walkable[K, V] {
	return &Walkable[K, V]{
		source: NewWalkableMap(items),
	}
}
