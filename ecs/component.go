package ecs

type ECSComponent interface {
    Id, EntityID uint64
    TypeId uint16
    Data map[string]interface{}
    Set(node string, data interface{})
    Delete(node string)
}

type Component struct {
	Id, EntityID uint64
	TypeId       uint16
	Data         map[string]interface{}
}

func (c *Component) ID() uint64 {
	return c.Id
}

func (c *Component) EntityID() uint64 {
	return c.EntityID
}

func (c *Component) TypeID() uint16 {
	return c.TypeId
}

func (c *Component) Set(node string, data interface{}) {
	c.Data[node] = data
}

func (c *Component) Delete(node string) {
	if n, ok := c.Data[node]; ok {
		delete(c.Data, node)
	}
}
