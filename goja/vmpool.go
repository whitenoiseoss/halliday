package goja

import (
	"sync"

	"github.com/dop251/goja"
)

// VMPool provides a pre-generated pool of prepped JavaScript runtimes
// for the scripting engine
type VMPool struct {
	lock sync.Mutex
	Pool []*goja.Runtime
}

// Get returns a *Runtime from the pool or generates a new one if the
// pool is empty
func (vmp *VMPool) Get() *goja.Runtime {
	vmp.lock.Lock()
	defer vmp.lock.Unlock()
	num := len(vmp.Pool)
	if num == 0 {
		return vmp.New()
	}

	// pop first element off of vmp.Pool
	r := vmp.Pool[0]
	copy(vmp.Pool, vmp.Pool[1:])
	vmp.Pool = vmp.Pool[:len(vmp.Pool)-1]
	return r
}

// New generates a new *Runtime
func (vmp *VMPool) New() *goja.Runtime {
	r := goja.New()
	// set the runtime up here
	// load scripts, set up global variables, new native funcs, etc
	return r
}

// Put places a *Runtime back into the VMPool
// Call this immediately after you finish with a runtime
func (vmp *VMPool) Put(r *goja.Runtime) {
	vmp.lock.Lock()
	defer vmp.lock.Unlock()
	vmp.Pool = append(vmp.Pool, r)
}
