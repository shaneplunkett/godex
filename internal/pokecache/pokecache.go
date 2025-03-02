package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Item map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Item[key] = cacheEntry{CreatedAt: time.Now(), Val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.Item[key]
	return entry.Val, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.Item {
			if now.Sub(entry.CreatedAt) > interval {
				delete(c.Item, key)
			}
		}

		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Item: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}
