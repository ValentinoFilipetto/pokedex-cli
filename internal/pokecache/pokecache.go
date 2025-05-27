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
	return &Cache{
		items:    make(map[string]cacheEntry),
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.items[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.items[key]
	if ok {
		return value.val, true
	}

	return nil, false
}
