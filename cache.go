package cache

import "sync"

type Cache interface {
	Set(key string, value any)
	Get(key string) any
	Delete(key string)
}

type InMemoryCache struct {
	cache map[string]any
	mu    sync.RWMutex
}

func New() InMemoryCache {
	return InMemoryCache{
		cache: make(map[string]any),
	}
}

func (i *InMemoryCache) Set(key string, value any) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.cache[key] = value
}

func (i *InMemoryCache) Get(key string) any {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.cache[key]
}

func (i *InMemoryCache) Delete(key string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	delete(i.cache, key)
}
