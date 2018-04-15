package ecs

type Entity struct {
	Id               uint64
	TypeId           uint16
	ComponentManager *ComponentManager
}

func (e *Entity) ID() uint64 {
	return e.Id
}

func (e *Entity) TypeID() uint16 {
	return e.TypeId
}
