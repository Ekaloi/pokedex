package pokecache

import (
	"fmt"
	"sync"
	"time"

	"github.com/Ekaloi/pokedex/poketypes"
)

type Cache struct {
	mu    sync.Mutex
	entry map[string]cacheEntry
	dex   map[string]poketypes.PokemonActual
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) *Cache {
	c := &Cache{
		entry: make(map[string]cacheEntry),
		dex:   make(map[string]poketypes.PokemonActual),
	}

	go c.reapLoop(t)

	return c
}

func (c *Cache) Add(key string, v []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{createdAt: time.Now(), val: v}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.entry[key]
	c.mu.Unlock()

	if !ok {
		return make([]byte, 0), ok
	}

	return entry.val, ok

}

func (c *Cache) Catch(name string, pokemondData poketypes.PokemonActual) {
	c.mu.Lock()
	c.dex[name] = pokemondData
	c.mu.Unlock()
}

func (c *Cache) Inspect(name string) (poketypes.PokemonActual, error) {
	c.mu.Lock()
	data, ok := c.dex[name]
	if !ok {
		return poketypes.PokemonActual{}, fmt.Errorf("you have not caught that pokemon")
	}
	c.mu.Unlock()

	return data, nil
}

func (c *Cache) InspectPokedex() ([]poketypes.PokemonActual, error) {
	res := make([]poketypes.PokemonActual, len(c.dex))
	c.mu.Lock()
	count := 0
	for _, v := range c.dex {
		res[count] = v
		count += 1
	}
	c.mu.Unlock()

	return res, nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		for k, v := range c.entry {
			dur := time.Since(v.createdAt)
			if dur > interval {
				c.mu.Lock()
				delete(c.entry, k)
				c.mu.Unlock()
			}
		}
	}
}
