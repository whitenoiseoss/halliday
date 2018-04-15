package ecs

type Component struct {
	Id       uint64
	Category string
}

func (c *Component) ID() uint64 {
	return c.Id
}
