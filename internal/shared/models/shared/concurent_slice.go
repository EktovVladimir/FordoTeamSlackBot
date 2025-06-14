package shared

import "sync"

type ConcurrentSlice[T any] struct {
	mu    sync.RWMutex
	slice []T
}

func NewConcurrentSlice[T any]() *ConcurrentSlice[T] {
	return &ConcurrentSlice[T]{
		mu:    sync.RWMutex{},
		slice: make([]T, 0),
	}
}

func (cs *ConcurrentSlice[T]) Append(ev T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.slice = append(cs.slice, ev)
}

func (cs *ConcurrentSlice[T]) Len() int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return len(cs.slice)
}

func (cs *ConcurrentSlice[T]) Slice(start int, end int) []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.slice[start:end]
}
