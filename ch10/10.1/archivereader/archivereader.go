package archivereader

import (
	"io"
	"sync"
)

var (
	registeredReader = make(map[string]func(string) (io.Reader, error))
	lock             sync.RWMutex
)

// Adaptor returns a new reader whereby data is decompressed
type Adaptor func(name string) (io.Reader, error)

// Register makes an adaptor available by the provided name
// Panics if the same name inserted twice or creator is nil
func Register(name string, f Adaptor) {
	lock.RLock()
	defer lock.Unlock()
	if f == nil {
		panic("archivereader: creator is nil")
	}
	if _, exists := registeredReader[name]; exists {
		panic("archivereader: Register called twice for " + name)
	}
	registeredReader[name] = f
}
