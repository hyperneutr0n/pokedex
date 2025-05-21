package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries				map[string]cacheEntry
	mu						*sync.RWMutex
	interval			time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, exist := c.entries[key]
	return data.val, exist
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.entries {
			if time.Since(val.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	createdAt		time.Time
	val					[]byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		mu: &sync.RWMutex{},
		interval: interval,
	}

	go c.reapLoop()
	return c
}