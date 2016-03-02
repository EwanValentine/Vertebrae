package vertebrae

import (
	"sync"
)

// Vertebrae is the foundation of the container store
// Please edit this to your liking, if you install a linter
// you will see that all exported items should be commented
type Vertebrae struct {
	// This is a locking mechanism. Locking your data in some way shape or form
	// is essential if your service is going to be used by concurrent processes.
	mux sync.RWMutex

	m map[string]interface{}
}

// Add takes a name and an object and adds it to the internal store
// Note: You want to make sure that you reference a pointer rather than a
// raw struct in a situation like this.
func (v *Vertebrae) Add(name string, object interface{}) {

	// Check if map has been initialised yet
	if v.m == nil {
		v.m = make(map[string]interface{})
	}

	v.mux.Lock()
	v.m[name] = object
	v.mux.Unlock()
}

// Remove service from internal store
func (v *Vertebrae) Remove(name string) {
	v.mux.Lock()
	delete(v.m, name)
	v.mux.Unlock()
}

// Get returns a service by name
func (v *Vertebrae) Get(name string) (object interface{}, ok bool) {
	v.mux.RLock()
	object, ok = v.m[name]
	v.mux.RUnlock()
	return object, ok
}
