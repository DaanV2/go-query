package query

import (
	"errors"
)

var (
	ErrNoMatches         = errors.New("no matches found")
	// ErrStopWalking is returned by the walk function to stop walking. And should be filtered out by the walk function.
	ErrStopWalking       = errors.New("stop walking")
	ErrWrongWalkableType = errors.New("wrong walkable type, cannot cast to IWalkable")
)

var _ IWalkable[int, int] = &WalkableError[int, int]{}
var _ error = &WalkableError[int, int]{}

type WalkableError[K any, V any] struct {
	err error
}

// Error implements error
func (w *WalkableError[K, V]) Error() string {
	return w.err.Error()
}

func NewWalkError[K any, V any](err error) *WalkableError[K, V] {
	return &WalkableError[K, V]{
		err: err,
	}
}

// Walk implements IWalkable
func (w *WalkableError[K, V]) Walk(call func(item V, key K) error) (IWalkable[K, V], error) {
	return w, w.err
}
