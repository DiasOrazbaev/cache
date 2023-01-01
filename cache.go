package cache

type Cache interface {
	Set(key string, value any)
	Get(key string) any
	Delete(key string)
}

type InMemoryCache struct {
	cache map[string]any
}

func New() InMemoryCache {
	return InMemoryCache{
		cache: make(map[string]any),
	}
}

func (i *InMemoryCache) Set(key string, value any) {
	i.cache[key] = value
}

func (i *InMemoryCache) Get(key string) any {
	return i.cache[key]
}

func (i *InMemoryCache) Delete(key string) {
	delete(i.cache, key)
}
