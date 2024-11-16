package cache

import "sync"

type Cache struct {
	store map[string]float64
	mu    sync.RWMutex
}

var tempCache = Cache{
	store: make(map[string]float64),
}

func Get(key string) (float64, bool) {
	tempCache.mu.RLock()
	defer tempCache.mu.RUnlock()
	value, found := tempCache.store[key]
	return value, found
}

func Set(key string, value float64) {
	tempCache.mu.Lock()
	defer tempCache.mu.Unlock()
	tempCache.store[key] = value
}