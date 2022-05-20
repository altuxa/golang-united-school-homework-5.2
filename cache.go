package cache

import (
	"time"
)

type Cache struct {
	data    map[string]string
	isExist map[string]time.Time
}

func NewCache() Cache {
	return Cache{
		data:    make(map[string]string),
		isExist: make(map[string]time.Time),
	}
}

func (c *Cache) IsExpire(key string, deadline time.Time) bool {
	if ok := time.Now().After(deadline); ok {
		return true
	}
	return false
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
	delete(c.isExist, key)
}

func (c *Cache) Get(key string) (string, bool) {
	if data, ok := c.data[key]; ok {
		if c.IsExpire(key, c.isExist[key]) {
			c.Delete(key)
		} else {
			return data, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	time := time.Date(3000, 1, 0, 0, 0, 0, 0, time.UTC)
	c.data[key] = value
	c.isExist[key] = time
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for key, _ := range c.data {
		if _, ok := c.data[key]; ok {
			deadline := c.isExist[key]
			if c.IsExpire(key, deadline) {
				c.Delete(key)
			} else {
				keys = append(keys, key)
			}
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.isExist[key] = deadline
}
