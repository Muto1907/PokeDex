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
	Entries   map[string]cacheEntry
	Protected *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{make(map[string]cacheEntry), &sync.Mutex{}}
	go cache.ReapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Protected.Lock()
	defer c.Protected.Unlock()
	entry := cacheEntry{time.Now(), val}
	c.Entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Protected.Lock()
	defer c.Protected.Unlock()
	val, ok := c.Entries[key]
	return val.val, ok
}

func (c *Cache) ReapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		c.Protected.Lock()
		for key, val := range c.Entries {
			if time.Since(val.createdAt) >= interval {
				delete(c.Entries, key)
			}
		}
		c.Protected.Unlock()
	}
}
