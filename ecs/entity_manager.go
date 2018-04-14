package ecs

import (
	"sync"
	"sync/atomic"
)

var (
	counterLock *sync.RWMutex
	entityInc   uint64
)

type EntityContainer map[uint64]*Entity

type EntityManager struct {
	Entities EntityContainer
}

func (em *EntityManager) CreateEntity() uint64 {
	ent := &Entity{id: em.NextEntityID(), typeid: 0}
	entid := ent.ID()
	em.Entities[entid] = ent
	return entid
}

// func (em *EntityManager) CreateEntities(amount int) []EntityID {
// }

// func (em *EntityManager) DestroyEntity(id EntityID) (something) {
// }

func (em *EntityManager) NextEntityID() uint64 {
	return atomic.AddUint64(&entityInc, 1)
}

func (em *EntityManager) Len() int {
	return len(em.Entities)
}
