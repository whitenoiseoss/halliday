package ecs

import "sync"

type ComponentID uint32

type Component struct {
	Id       ComponentID
	datalock *sync.RWMutex
	Data     map[uint64]interface{}
}

func (c *Component) ID() ComponentID {
	return c.Id
}
