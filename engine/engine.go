package engine

import (
	"github.com/whitenoiseoss/halliday/ecs"
	"github.com/whitenoiseoss/halliday/goja"
)

// HallidayEngine is a wrapper around all other Engine
// or Manager objects. It represents a single source of truth
// for the entire game state
type HallidayEngine struct {
	ECS       *ecs.ECSEngine
	VMManager *goja.VMManager
}

// NewHallidayEngine creates a new HallidayEngine instance
func NewHallidayEngine() *HallidayEngine {
	he := &HallidayEngine{
		ECS:       ecs.NewECSEngine(),
		VMManager: goja.NewVMManager(),
	}

	return he
}
