package ecs

import (
	"sort"
	"sync"
	"sync/atomic"
)

type EntityManager struct {
	sync.RWMutex
	inc              uint64
	Entities         map[uint64]*Entity
	ComponentManager *ComponentManager
}

func NewEntityManager(cm *ComponentManager) *EntityManager {
	em := new(EntityManager)
	em.Entities = make(map[uint64]*Entity)
	em.ComponentManager = cm

	return em
}

// Creates an Entity and indexes it by ID (uint64)
func (em *EntityManager) CreateEntity() uint64 {
	// TODO: clean this up
	ent := &Entity{Id: em.NextEntityID(), TypeId: 0}
	// pass the ComponentManager along
	ent.ComponentManager = em.ComponentManager
	entid := ent.ID()
	em.Lock()
	em.Entities[entid] = ent
	em.Unlock()

	return entid
}

// Create Entity objects in bulk
// func (em *EntityManager) CreateEntities(amount int) []uint64 {
//
// }

// TODO: Create full destruction process to recycle ID, release components, etc
func (em *EntityManager) DestroyEntity(id uint64) {
	delete(em.Entities, id)
}

// TODO: double return (*Entity, error) ?
// Return an Entity or nil, searching by ID
func (em *EntityManager) GetEntityByID(id uint64) *Entity {
	em.RLock()
	if entity, ok := em.Entities[id]; ok {
		return entity
	}
	em.RUnlock()

	return nil
}

// Returns a sorted list of all Entities
func (em *EntityManager) All() []*Entity {
	var keys IDList
	var entities []*Entity

	for key := range em.Entities {
		keys = append(keys, key)
	}

	sort.Sort(keys)
	for _, key := range keys {
		entities = append(entities, em.Entities[key])
	}

	return entities
}

// Returns next valid Entity ID, starts at 0
func (em *EntityManager) NextEntityID() uint64 {
	// start at 0
	return atomic.AddUint64(&em.inc, 1) - 1
}

// Returns number of Entities
func (em *EntityManager) Len() int {
	return len(em.Entities)
}
