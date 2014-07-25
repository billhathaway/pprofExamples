// badLRU project badLRU.go
package badLRU

import "sync"
import "time"

type Cache struct {
	sizeLimit int
	data      map[int]*cacheEntry
	sync.RWMutex
}

type cacheEntry struct {
	value int
	time  time.Time
}

func New(size int) *Cache {
	return &Cache{sizeLimit: size, data: make(map[int]*cacheEntry)}
}

// horribly inefficient to allow us to show profiler
// not calling Lock() since we are already locked in Put()
func (c *Cache) makeRoom() {
	var oldestKey int
	oldestTime := time.Now()
	for key, value := range c.data {
		if value.time.Before(oldestTime) {
			oldestKey = key
			oldestTime = value.time
		}
	}
	delete(c.data, oldestKey)
}

func (c *Cache) Put(key int, val int) {
	c.Lock()
	defer c.Unlock()
	if len(c.data) >= c.sizeLimit {
		c.makeRoom()
	}
	c.data[key] = &cacheEntry{value: val}
}

func (c *Cache) Get(key int) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	entry, exists := c.data[key]
	if exists {
		c.data[key].time = time.Now()
		return entry.value, exists
	}
	return 0, exists
}
