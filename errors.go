package query

import (
	"errors"
)

var (
	// ErrNoMatches is returned by when no matches are found.
	ErrNoMatches = errors.New("no matches found")
	// ErrStopWalking is returned to stop walking. And should be filtered out by the walk function.
	ErrStopWalking = errors.New("stop walking")
	// ErrWrongWalkableType is returned if the type is not walkable.
	ErrWrongWalkableType = errors.New("wrong walkable type, cannot cast to IWalkable")
)

var _ IWalkable[int, int] = &WalkableError[int, int]{}
var _ error = &WalkableError[int, int]{}

// WalkableError is a walker that always returns an error.
type WalkableError[K any, V any] struct {
	err error
}

// Error implements error
func (w *WalkableError[K, V]) Error() string {
	return w.err.Error()
}

// NewWalkError creates a new walkable error.
func NewWalkError[K any, V any](err error) *WalkableError[K, V] {
	return &WalkableError[K, V]{
		err: err,
	}
}

// Walk implements IWalkable
func (w *WalkableError[K, V]) Walk(call func(item V, key K) error) (IWalkable[K, V], error) {
	return w, w.err
}
