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

func (cr *ComponentRegistry) RegisterComponent() {
}
