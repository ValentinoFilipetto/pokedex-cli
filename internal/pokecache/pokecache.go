package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	items    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:       sync.Mutex{},
		items:    make(map[string]cacheEntry),
		interval: interval,
	}
	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	c.items[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.items[key]

	if ok {
		return value.val, true
	}

	return nil, false
}

// Each time c.Interval passes we should remove entries that are older than c.Interval
// This ensures that the cache does not grow too large over time.
func (c *Cache) reapLoop() {
	go func() {
		ticker := time.NewTicker(c.interval)
		for {
			<-ticker.C
			c.mu.Lock()
			now := time.Now()
			for key, item := range c.items {
				if item.createdAt.Add(c.interval).Before(now) {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
