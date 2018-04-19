package ecs

import (
	"sync"
	"sync/atomic"
)

// get component by id?
// get components by type
// get components on entity id

// ComponentManager is a struct factory
// possible that ComponentManager will go on Entity
type ComponentManager struct {
	sync.RWMutex
	inc               uint64
	ComponentRegistry *ComponentRegistry
}

func NewComponentManager(cr *ComponentRegistry) *ComponentManager {
	cm := new(ComponentManager)
	cm.ComponentRegistry = cr

	return cm
}

// func CreateBaseComponent()

// func CreateComponent() -- read from ComponentRegistry

// func AddComponent(ent uint64)

// Returns next valid Component ID, starts at 0
func (cm *ComponentManager) NextComponentID() uint64 {
	// start at 0
	return atomic.AddUint64(&cm.inc, 1) - 1
}
