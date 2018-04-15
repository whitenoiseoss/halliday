package ecs

import "sync"

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
