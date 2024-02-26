package tests

import (
	"github.com/zeynepbetul/pokedexcli/internal/pokecache"
	"testing"
	"time"
)

// Test Reap Loop
// Reap Loop deletes entries which entries are higher than an interval
func TestReapLoop(t *testing.T) {
	const interval = 35 * time.Second

	TestEntry := pokecache.CacheEntry{
		CreatedAt: time.Now(),
		Val:       []byte{78, 111, 116, 32, 70, 111, 117, 110, 100},
	}
	TestEntryToDelete := pokecache.CacheEntry{
		CreatedAt: time.Now().Add(35 * time.Second),
		Val:       []byte{78, 111, 116, 32, 70, 111, 117, 110, 100},
	}

	TestCacheEntries := make(map[string]pokecache.CacheEntry)

	TestCacheEntries["https://pokeapi.co/api/v2/location-area/"] = TestEntry
	TestCacheEntries["https://pokeapi.co/api/v2/location-area/asd"] = TestEntryToDelete

	TestCache := pokecache.Cache{
		CacheEntries: TestCacheEntries,
		Interval:     interval,
	}
	TestCache.ReapLoop()
}
