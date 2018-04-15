package ecs

// ComponentManager is a struct factory
// possible that ComponentManager will go on Entity
type ComponentManager struct{}

func NewComponentManager() *ComponentManager {
	cm := new(ComponentManager)

	return cm
}
