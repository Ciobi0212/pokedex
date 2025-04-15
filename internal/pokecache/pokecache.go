package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	Entry     []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	stopChan     chan struct{}
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()

	defer c.mu.Unlock()

	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		Entry:     value,
	}

	c.cacheEntries[key] = cacheEntry
}

func (c *Cache) Get(key string) (cacheEntry, bool) {

	c.mu.Lock()

	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]

	return entry, ok
}

func (c *Cache) ReadLoop(interval time.Duration) {
	ticker := time.Tick(interval)

	for {
		select {
		case <-ticker:
			c.mu.Lock()

			for key, value := range c.cacheEntries {
				elapsedTimeSinceCreation := time.Since(value.createdAt)
				if elapsedTimeSinceCreation > interval {
					delete(c.cacheEntries, key)
				}
			}

			c.mu.Unlock()

		case <-c.stopChan:
			fmt.Println("Cache reaping goroutine stopped")
			return
		}
	}
}

func (c *Cache) StopReaping() {
	close(c.stopChan)
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: map[string]cacheEntry{},
		mu:           sync.Mutex{},
		stopChan:     make(chan struct{}),
	}

	go cache.ReadLoop(interval)

	return cache
}
