package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachemap map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	m := make(map[string]cacheEntry)

	c := Cache{
		cachemap: m,
		interval: interval,
	}
	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) error {
	ce := cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Lock()
	c.cachemap[key] = ce
	c.mu.Unlock()
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	ce, ok := c.cachemap[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return ce.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for k, entry := range c.cachemap {
			age := time.Since(entry.createdAt)
			if age > c.interval {
				delete(c.cachemap, k)
			}
		}
		c.mu.Unlock()
	}
}
