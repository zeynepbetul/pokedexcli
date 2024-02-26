package pokecache

import (
	"sync"
	"time"
)

var interval = 35 * time.Second
var FirstCache = NewCache(interval)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte // represents the raw data we're caching
}
type Cache struct {
	CacheEntries map[string]CacheEntry
	Mu           sync.Mutex
	Interval     time.Duration
}

func NewCache(interval time.Duration) Cache {
	cacheEntries := make(map[string]CacheEntry)

	newC := Cache{
		CacheEntries: cacheEntries,
		Interval:     interval,
	}
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			newC.ReapLoop()
		}
	}()
	return newC
}

func (c *Cache) ReapLoop() {
	for key, value := range c.CacheEntries {
		if time.Now().Sub(value.CreatedAt) > c.Interval {
			c.Mu.Lock()
			delete(c.CacheEntries, key)
			c.Mu.Unlock()
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.CacheEntries[key] = CacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.CacheEntries[key]
	return cacheEntry.Val, ok
}
