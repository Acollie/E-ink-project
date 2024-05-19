package cache

import "time"

func (c *Cache) isValid() bool {
	return c.lastUpdated+c.cacheExpire > time.Now().Unix()
}

func (c *Cache) Fetch() *interface{} {
	if c.isValid() {
		return &c.Result
	}
	return nil
}

func (c *Cache) Update(response interface{}) {
	c.lastUpdated = time.Now().Unix()
	c.Result = response
}
