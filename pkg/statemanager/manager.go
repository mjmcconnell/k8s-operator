package statemanager

import (
	"sync"
)

var (
	manager *StateManager
	once    sync.Once
)

// StateManager defines a provides a structure for storing global state of resources within the cluster
// it uses the sync.Mutex library, to providing locking
type StateManager struct {
	*sync.Mutex
	state map[string]interface{}
}

// GetStateManager returns an instance of the state manager
func GetStateManager() *StateManager {
	once.Do(func() {
		manager = &StateManager{Mutex: &sync.Mutex{}}
		manager.state = make(map[string]interface{})
	})
	return manager
}

// Get looks up the state manager for existing value
func (sm *StateManager) Get(key string) interface{} {
	sm.Lock()
	defer sm.Unlock()
	return sm.state[key]
}

// Set creates a new value in the state manager
func (sm *StateManager) Set(key string, value interface{}) {
	sm.Lock()
	defer sm.Unlock()
	sm.state[key] = value
}

// Clear resets the state manager
func (sm *StateManager) Clear() {
	sm.Lock()
	defer sm.Unlock()
	sm.state = make(map[string]interface{})
}
