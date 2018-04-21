package main

import (
	"fmt"
	"reflect"

	"github.com/dop251/goja"
	"github.com/whitenoiseoss/halliday/ecs"
)

// Test struct is for testing ECSComponent interface
type Test struct {
	*ecs.Component
	Name string
}

// Bark is a method that makes Test act like a dog
func (t *Test) Bark() {
	fmt.Println("Woof")
}

func main() {
	vm := goja.New()
	ECSEngine := ecs.NewECSEngine()
	eid := ECSEngine.EntityManager.CreateEntity()
	eid2 := ECSEngine.EntityManager.CreateEntity()
	fmt.Println(eid)
	fmt.Println(eid2)
	fmt.Println(ECSEngine.EntityManager.Len())
	ECSEngine.ComponentRegistry.RegisterComponent("Test", Test{})
	T := ECSEngine.ComponentRegistry.Type("Test")
	fmt.Println(T)
	testObj := reflect.New(T).Interface().(*Test)
	testObj.Bark()

	vm.Set("ECS", ECSEngine)
	vm.Set("EntityManager", ECSEngine.EntityManager)
	js, _ := vm.RunString(`
	var eid = EntityManager.CreateEntity();
	EntityManager.GetEntityByID(eid);
	`)

	eid3 := js.Export().(*ecs.Entity)
	fmt.Println(eid3)
	fmt.Println(ECSEngine.EntityManager.Len())
}
