package workitem

import (
	"sync"

	"github.com/almighty/almighty-core/log"
)

// WorkItemTypeCache represents WorkItemType cache
type WorkItemTypeCache struct {
	cache   map[string]WorkItemType
	mapLock sync.RWMutex
}

// NewWorkItemTypeCache constructs WorkItemTypeCache
func NewWorkItemTypeCache() *WorkItemTypeCache {
	witCache := WorkItemTypeCache{}
	witCache.cache = make(map[string]WorkItemType)
	return &witCache
}

// Get returns WorkItemType by name.
// The second value (ok) is a bool that is true if the WorkItemType exists in the cache, and false if not.
func (c *WorkItemTypeCache) Get(typeName string) (WorkItemType, bool) {
	c.mapLock.RLock()
	defer c.mapLock.RUnlock()
	w, ok := c.cache[typeName]
	return w, ok
}

// Put puts a work item type to the cache
func (c *WorkItemTypeCache) Put(wit WorkItemType) {
	c.mapLock.Lock()
	defer c.mapLock.Unlock()
	c.cache[wit.Name] = wit
}

// Clear clears the cache
func (c *WorkItemTypeCache) Clear() {
	c.mapLock.Lock()
	defer c.mapLock.Unlock()
	log.Logger().Infoln("Clearing work item cache")

	c.cache = make(map[string]WorkItemType)
}
