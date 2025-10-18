package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu		   sync.Mutex	
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time 
	val				[]byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.entries[key]
	if !ok {
		return nil, false 
	}
	return e.val, true 
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	
	for range ticker.C {
		cutoff := time.Now().Add(-c.interval)


		c.mu.Lock()
		for k, entry := range c.entries {
			if entry.createdAt.Before(cutoff) {
					delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}

	go c.reapLoop() 
	return c
}
