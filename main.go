package main

import (
	"fmt"

	"github.com/whitenoiseoss/halliday/ecs"
)

func main() {
	ECS_Engine := ecs.NewEngine()
	eid := ECS_Engine.EntityManager.CreateEntity()
	fmt.Println(eid)
	fmt.Println(ECS_Engine.EntityManager.Len())
}
