package db

import "sync"

type entityContext[T any] struct {
	mu    *sync.RWMutex
	slice []*T
	maxId UniqId
}

func NewEntityContext[T any]() *entityContext[T] {
	return &entityContext[T]{
		mu:    &sync.RWMutex{},
		slice: make([]*T, 0),
		maxId: 0,
	}
}

func (ec *entityContext[T]) GetWithLock() []*T {
	ec.mu.RLock()
	defer ec.mu.RUnlock()
	return ec.slice
}

func (ec *entityContext[T]) Append(item *T) {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	ec.slice = append(ec.slice, item)
}
