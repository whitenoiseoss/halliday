package ecs

import (
	"sync"
	"sync/atomic"
)

// ComponentManager is a struct factory
// possible that ComponentManager will go on Entity
type ComponentManager struct {
	sync.RWMutex
	inc uint64
}

func NewComponentManager() *ComponentManager {
	cm := new(ComponentManager)

	return cm
}

// func CreateBaseComponent()

// func CreateComponent() -- read from ComponentRegistry

// Returns next valid Component ID, starts at 0
func (cm *ComponentManager) NextComponentID() uint64 {
	// start at 0
	return atomic.AddUint64(&cm.inc, 1) - 1
}
