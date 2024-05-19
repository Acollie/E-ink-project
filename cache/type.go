package cache

type Cache struct {
	lastUpdated int64
	cacheExpire int64
	Result      interface{}
}

func New(cacheExpire int64, response interface{}) *Cache {
	return &Cache{
		lastUpdated: 0,
		cacheExpire: cacheExpire,
		Result:      response,
	}
}
