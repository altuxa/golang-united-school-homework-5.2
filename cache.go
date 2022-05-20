package cache

import "time"

type Cache struct {
	data map[string]string
}

func NewCache() Cache {
	return Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	if data, ok := c.data[key]; ok {
		return data, true
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
}

func (c Cache) Keys() []string {
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
}
