package ecs

type Engine struct {
	EntityManager *EntityManager
	// ComponentManager *ComponentManager
	// SystemManager    *SystemManager
}

func NewEngine() *Engine {
	engine := &Engine{EntityManager: &EntityManager{Entities: make(EntityContainer)}}
	// ComponentManager: &ComponentManager{},
	// SystemManager: &SystemManager{}
	return engine
}
