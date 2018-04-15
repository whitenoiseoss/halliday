package main

import (
	"fmt"
	"reflect"

	"github.com/whitenoiseoss/halliday/ecs"
)

type Test struct {
	*ecs.Component
	Name string
}

func (t *Test) Bark() {
	fmt.Println("Woof")
}

func main() {
	ECS_Engine := ecs.NewECSEngine()
	eid := ECS_Engine.EntityManager.CreateEntity()
	eid2 := ECS_Engine.EntityManager.CreateEntity()
	fmt.Println(eid)
	fmt.Println(eid2)
	fmt.Println(ECS_Engine.EntityManager.Len())
	ECS_Engine.ComponentRegistry.RegisterComponent("Test", Test{})
	T := ECS_Engine.ComponentRegistry.Type("Test")
	fmt.Println(T)
	testObj := reflect.New(T).Interface().(*Test)
	testObj.Bark()
}
