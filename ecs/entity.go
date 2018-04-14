package ecs

type Entity struct {
	id     uint64
	typeid uint16
}

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) Type() uint16 {
	return e.typeid
}
