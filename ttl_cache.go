package cache

import (
	"sync"
	"time"
)

type TTLCache interface {
	Set(key string, value any, t time.Duration)
	Get(key string) (any, error)
	Delete(key string) error
}

type cache struct {
	data     any
	tStart   time.Time
	duration time.Duration
}

type TTLInMemoryCache struct {
	mu    sync.RWMutex
	cache map[string]cache
}

func (T *TTLInMemoryCache) scheduling() {
	for {
		T.mu.Lock()
		for k, v := range T.cache {
			if time.Since(v.tStart) > v.duration {
				delete(T.cache, k)
			}
		}
		T.mu.Unlock()
	}
}

func NewTTLInMemoryCache() *TTLInMemoryCache {
	l := &TTLInMemoryCache{
		cache: make(map[string]cache),
	}
	go l.scheduling()
	return l
}

func (T *TTLInMemoryCache) Set(key string, value any, t time.Duration) {
	T.mu.Lock()
	defer T.mu.Unlock()
	T.cache[key] = cache{value, time.Now(), t}
}

func (T *TTLInMemoryCache) Get(key string) (any, error) {
	T.mu.RLock()
	defer T.mu.RUnlock()
	d, s := T.cache[key]
	if !s {
		return nil, ErrNotFound
	}
	return d.data, nil
}

func (T *TTLInMemoryCache) Delete(key string) error {
	T.mu.Lock()
	defer T.mu.Unlock()
	_, s := T.cache[key]
	if !s {
		return ErrNotFound
	}
	delete(T.cache, key)
	return nil
}
