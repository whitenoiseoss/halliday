package main

import (
	"fmt"
	"reflect"

	"github.com/whitenoiseoss/halliday/ecs"
	"github.com/whitenoiseoss/halliday/engine"
)

// Test struct is for testing ECSComponent interface
type Test struct {
	*ecs.Component
	Name string
}

// NewTest returns a new Test object
func NewTest() ecs.ECSComponent {
	return Test{}
}

// Bark is a method that makes Test act like a dog
func (t *Test) Bark() {
	fmt.Println("Woof")
}

func main() {
	Halliday := engine.NewHallidayEngine()
	eid := Halliday.ECS.EntityManager.CreateEntity()
	fmt.Println(eid)

	// Halliday.ECS.ComponentRegistry.RegisterComponent("Test", Test{})

	vm := Halliday.VMManager.VMPool.Get()
	js, _ := vm.RunString(`
		2+2
	`)
	result := js.Export().(int64)
	fmt.Println(result)

	vm.Set("Test", Test{})
	vm.Set("NewTest", NewTest)
	vm.Set("ComponentRegistry", Halliday.ECS.ComponentRegistry)

	_, _ = vm.RunString(`
		ComponentRegistry.RegisterComponent("Test", NewTest());
	`)

	Halliday.VMManager.ReturnVM(vm)
	fmt.Println(Halliday.VMManager.VMPool.Len())
	vm = Halliday.VMManager.GetVM()
	fmt.Println(Halliday.VMManager.VMPool.Len())

	T := Halliday.ECS.ComponentRegistry.Type("Test")
	testObj := reflect.New(T).Interface().(*Test)
	testObj.Bark()

	fmt.Println(Halliday.ECS.EntityManager.Len())
	// vm := goja.New()
	// ECSEngine := ecs.NewECSEngine()
	// eid := ECSEngine.EntityManager.CreateEntity()
	// eid2 := ECSEngine.EntityManager.CreateEntity()
	// fmt.Println(eid)
	// fmt.Println(eid2)
	// fmt.Println(ECSEngine.EntityManager.Len())
	// ECSEngine.ComponentRegistry.RegisterComponent("Test", Test{})
	// T := ECSEngine.ComponentRegistry.Type("Test")
	// fmt.Println(T)
	// testObj := reflect.New(T).Interface().(*Test)
	// testObj.Bark()

	// vm.Set("ECS", ECSEngine)
	// vm.Set("EntityManager", ECSEngine.EntityManager)
	// js, _ := vm.RunString(`
	// var eid = EntityManager.CreateEntity();
	// EntityManager.GetEntityByID(eid);
	// `)

	// eid3 := js.Export().(*ecs.Entity)
	// fmt.Println(eid3)
	// fmt.Println(ECSEngine.EntityManager.Len())
}
