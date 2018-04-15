package ecs

type ECSComponent interface {
	Set(node string, data interface{})
	Delete(node string)
}

type Component struct {
	Id       uint64
	EntityID uint64
	TypeId   uint16
	Data     map[string]interface{}
	Enabled  bool
}

func (c *Component) ID() uint64 {
	return c.Id
}

func (c *Component) Owner() uint64 {
	return c.EntityID
}

func (c *Component) TypeID() uint16 {
	return c.TypeId
}

func (c *Component) Set(node string, data interface{}) {
	c.Data[node] = data
}

func (c *Component) Delete(node string) {
	delete(c.Data, node)
}
