package backup

import (
	"sync"
)

// backedUpItemsTracker keeps track of items which have been backed up
type backedUpItemsTracker struct {
	*sync.RWMutex
	items map[itemKey]struct{}
}

func NewBackedUpItemsTracker() *backedUpItemsTracker {
	return &backedUpItemsTracker{
		RWMutex: &sync.RWMutex{},
		items:   map[itemKey]struct{}{},
	}
}

func (bit *backedUpItemsTracker) Add(key itemKey) {
	bit.Lock()
	defer bit.Unlock()
	bit.items[key] = struct{}{}

}

// Count of backed up items
func (bit *backedUpItemsTracker) Count() int {
	bit.Lock()
	defer bit.Unlock()
	return len(bit.items)
}

// Check if a given item is already backed up or not.
func (bit *backedUpItemsTracker) Exists(key itemKey) bool {
	bit.RLock()
	defer bit.RUnlock()
	_, ok := bit.items[key]
	return ok
}
