package ecs

type ECSEngine struct {
	EntityManager     *EntityManager
	ComponentManager  *ComponentManager
	ComponentRegistry *ComponentRegistry
	SystemManager     *SystemManager
}

func NewECSEngine() *ECSEngine {
	engine := new(ECSEngine)
	engine.ComponentRegistry = NewComponentRegistry()
	engine.ComponentManager = NewComponentManager(engine.ComponentRegistry)
	engine.EntityManager = NewEntityManager(engine.ComponentManager)
	engine.SystemManager = NewSystemManager()

	return engine
}
