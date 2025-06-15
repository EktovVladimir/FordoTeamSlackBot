package db_adapter

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"maps"
	"slices"
	"sync"
)

type UniqEntity interface {
	GetId() types.UniqId
	SetId(id types.UniqId)
}

type Entity[T UniqEntity] struct {
	mu    *sync.RWMutex
	items map[types.UniqId]T
	maxId types.UniqId
}

func NewEntity[T UniqEntity]() *Entity[T] {
	return &Entity[T]{
		mu:    &sync.RWMutex{},
		items: make(map[types.UniqId]T),
		maxId: 0,
	}
}

func NewLoadedEntity[T UniqEntity](slice []T) *Entity[T] {
	entity := NewEntity[T]()

	for _, e := range slice {
		entity.maxId = maxUniqId(entity.maxId, e.GetId())
		entity.items[e.GetId()] = e
	}

	return entity
}

func (ec *Entity[T]) GetAll() []T {
	ec.mu.RLock()
	defer ec.mu.RUnlock()

	if len(ec.items) == 0 {
		return make([]T, 0)
	}

	return slices.Collect(maps.Values(ec.items))
}

func (ec *Entity[T]) Get(id types.UniqId) (T, bool) {
	ec.mu.RLock()
	defer ec.mu.RUnlock()

	value, found := ec.items[id]
	return value, found
}

func (ec *Entity[T]) Insert(item T) {
	ec.mu.Lock()
	defer ec.mu.Unlock()

	ec.maxId = generateUniqId(ec.maxId)
	item.SetId(ec.maxId)
	ec.items[ec.maxId] = item
}

func (ec *Entity[T]) Update(item T) {
	ec.mu.Lock()
	defer ec.mu.Unlock()

	id := item.GetId()
	ec.items[id] = item
}

func (ec *Entity[T]) Delete(id types.UniqId) {
	ec.mu.Lock()
	defer ec.mu.Unlock()

	delete(ec.items, id)
}

func (ec *Entity[T]) BulkInsert(items []T) {
	ec.mu.Lock()
	defer ec.mu.Unlock()

	for _, item := range items {
		ec.maxId = generateUniqId(ec.maxId)
		item.SetId(ec.maxId)
		ec.items[ec.maxId] = item
	}
}

func generateUniqId(oldId types.UniqId) types.UniqId {
	return oldId + 1
}

func maxUniqId(a types.UniqId, b types.UniqId) types.UniqId {
	if a > b {
		return a
	}
	return b
}
