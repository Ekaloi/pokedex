package pokecache

import (
	"sync"
	"time"
)
type Cache struct {
	mu 	sync.Mutex
	entry map[string]cacheEntry
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(t time.Duration) *Cache{	
	c := &Cache{
		entry:     make(map[string]cacheEntry)}

	go c.reapLoop(t)

	return c
}

func (c *Cache) Add(key string, v []byte){
	c.mu.Lock()
	c.entry[key] = cacheEntry{createdAt: time.Now(), val: v}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry , ok := c.entry[key]
	c.mu.Unlock()

	if !ok {
		return make([]byte, 0), ok
	}

	return entry.val, ok

}

func (c *Cache) reapLoop(interval time.Duration){
	for{
		time.Sleep(interval)
		for k, v := range c.entry{
			dur := time.Since(v.createdAt)
			if dur > interval{
				c.mu.Lock()
				delete(c.entry, k)
				c.mu.Unlock()
			}
		}
	}
}