package pokecache

import (
	"sync"
	"time"
)

var interval = 35 * time.Second
var Cache = NewCache(interval)

type cacheEntry struct {
	createdAt time.Time
	val       []byte // represents the raw data we're caching
}
type cache struct {
	cacheEntries map[string]cacheEntry
	Mu           sync.Mutex
	interval     time.Duration
}

func NewCache(interval time.Duration) cache {
	cacheEntries := make(map[string]cacheEntry)

	newC := cache{
		cacheEntries: cacheEntries,
		interval:     interval,
	}
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			newC.reapLoop()
		}
	}()
	return newC
}

func (c *cache) reapLoop() {
	for key, value := range c.cacheEntries {
		if time.Now().Sub(value.createdAt) > c.interval {
			c.Mu.Lock()
			delete(c.cacheEntries, key)
			c.Mu.Unlock()
		}
	}
}

func (c *cache) Add(key string, val []byte) {
	c.cacheEntries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.cacheEntries[key]
	return cacheEntry.val, ok
}
