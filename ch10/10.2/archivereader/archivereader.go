package archivereader

import (
	"os"
	"sync"
)

var (
	registeredReader = make(map[string]Adaptor)
	lock             sync.RWMutex
)

type inforer interface {
	FileInfo() os.FileInfo
}

// Adaptor returns a new reader whereby data is decompressed
type Adaptor interface {
	Open(name string) ([]inforer, error)
}

// Register makes an adaptor available by the provided name
// Panics if the same name inserted twice or creator is nil
func Register(name string, a Adaptor) {
	lock.RLock()
	defer lock.Unlock()
	if a == nil {
		panic("archivereader: creator is nil")
	}
	if _, exists := registeredReader[name]; exists {
		panic("archivereader: Register called twice for " + name)
	}
	registeredReader[name] = a
}
