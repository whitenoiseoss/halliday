package ecs

type ECSEngine struct {
	EntityManager    *EntityManager
	ComponentManager *ComponentManager
	SystemManager    *SystemManager
}

func NewECSEngine() *ECSEngine {
	engine := new(ECSEngine)
	engine.ComponentManager = NewComponentManager()
	engine.EntityManager = NewEntityManager(engine.ComponentManager)
	engine.SystemManager = NewSystemManager()

	return engine
}
