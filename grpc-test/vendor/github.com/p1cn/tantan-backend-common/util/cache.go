package util

import (
	"sync"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

type Item struct {
	Object     interface{}
	Expiration int64
}

type Cache struct {
	mu    *sync.RWMutex
	items map[string]*Item
}

func NewCache() *Cache {
	return &Cache{
		mu:    &sync.RWMutex{},
		items: make(map[string]*Item),
	}
}

func (c *Cache) Size() int {
	return len(c.items)
}

func (c *Cache) Set(k string, x interface{}, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = NoExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	c.items[k] = &Item{
		Object:     x,
		Expiration: e,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, false
		}
	}
	c.mu.RUnlock()
	return item.Object, true
}
