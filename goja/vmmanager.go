package goja

import "github.com/dop251/goja"

// VMManager is a wrapper around VMPool and takes any helper functions
type VMManager struct {
	VMPool *VMPool
}

// NewVMManager returns a new instance of VMManager
func NewVMManager() *VMManager {
	vmm := &VMManager{
		VMPool: &VMPool{
			Pool: make([]*goja.Runtime, 0, 5),
		},
	}

	return vmm
}

// GetVM grabs a new VM from VMManager's VMPool
func (vmm *VMManager) GetVM() *goja.Runtime {
	return vmm.VMPool.Get()
}

// ReturnVM sends a *Runtime back to the VMPool
func (vmm *VMManager) ReturnVM(r *goja.Runtime) {
	// TODO: Maybe have this return a true/false/error
	vmm.VMPool.Put(r)
}
