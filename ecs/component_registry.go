package ecs

import (
	"reflect"
	"sync"
)

// build a reflect.TypeOf() Component registry
type ComponentRegistry struct {
	sync.RWMutex
	Catalog map[string]reflect.Type
}

func NewComponentRegistry() *ComponentRegistry {
	cr := new(ComponentRegistry)
	cr.Catalog = make(map[string]reflect.Type)

	return cr
}

func (cr *ComponentRegistry) RegisterComponent(name string, comp ECSComponent) error {
	if _, exists := cr.Catalog[name]; exists {
		// TODO: return a Halliday error
		return nil
	}

	cr.Catalog[name] = reflect.TypeOf(comp)
	return nil
}

func (cr *ComponentRegistry) Type(name string) reflect.Type {
	return cr.Catalog[name]
}
