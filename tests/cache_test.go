package tests

import (
	"github.com/zeynepbetul/pokedexcli/internal/pokecache"
	"testing"
	"time"
)

// Test Reap Loop
// Reap Loop deletes entries which entries are higher than an interval
func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Second

	TestEntry := pokecache.CacheEntry{
		CreatedAt: time.Now(),
		Val:       []byte{78, 111, 116, 32, 70, 111, 117, 110, 100},
	}

	TestCacheEntries := make(map[string]pokecache.CacheEntry)

	TestCacheEntries["https://pokeapi.co/api/v2/location-area/"] = TestEntry

	TestCache := pokecache.Cache{
		CacheEntries: TestCacheEntries,
		Interval:     interval,
	}
	_, ok := TestCache.Get("https://pokeapi.co/api/v2/location-area/")
	if !ok {
		t.Errorf("expected to find cache entry")
	}
	time.Sleep(interval)
	TestCache.ReapLoop() //
	_, ok = TestCache.Get("https://pokeapi.co/api/v2/location-area/")
	if ok {
		t.Errorf("expected to not find cache entry")
	}
}
