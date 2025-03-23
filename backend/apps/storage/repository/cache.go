package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	cacheDir = "/app/cache"
)

type CacheManager struct {
	mu  sync.RWMutex
	ttl time.Duration
}

type CacheItem struct {
	Value      any   `json:"value"`
	Expiration int64 `json:"expiration"`
}

func NewCacheManager(ttl time.Duration) *CacheManager {
	return &CacheManager{
		ttl: ttl,
	}
}

func (cm *CacheManager) Set(key string, value any) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	item := CacheItem{
		Value:      value,
		Expiration: time.Now().Add(cm.ttl).UnixNano(),
	}

	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal cache item: %v", err)
	}

	filePath := filepath.Join(cacheDir, key+".json")
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %v", err)
	}

	log.Printf("Cached item with key %s", key)
	return nil
}

func (cm *CacheManager) Get(key string, result any) (bool, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	filePath := filepath.Join(cacheDir, key+".json")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to read cache file: %v", err)
	}

	var item CacheItem
	if err := json.Unmarshal(data, &item); err != nil {
		return false, fmt.Errorf("failed to unmarshal cache item: %v", err)
	}

	if item.Expiration < time.Now().UnixNano() {
		os.Remove(filePath)
		return false, nil
	}

	valueData, err := json.Marshal(item.Value)
	if err != nil {
		return false, fmt.Errorf("failed to marshal item value: %v", err)
	}

	if err := json.Unmarshal(valueData, result); err != nil {
		return false, fmt.Errorf("failed to unmarshal into result: %v", err)
	}

	log.Printf("Retrieved cached item with key %s", key)
	return true, nil
}

func (cm *CacheManager) Delete(key string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	filePath := filepath.Join(cacheDir, key+".json")
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete cache file: %v", err)
	}

	log.Printf("Deleted cached item with key %s", key)
	return nil
}

func (cm *CacheManager) Clear() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	dir, err := os.Open(cacheDir)
	if err != nil {
		return fmt.Errorf("failed to open cache directory: %v", err)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return fmt.Errorf("failed to read cache directory: %v", err)
	}

	for _, name := range names {
		if filepath.Ext(name) == ".json" {
			filePath := filepath.Join(cacheDir, name)
			if err := os.Remove(filePath); err != nil {
				log.Printf("Warning: failed to delete cache file %s: %v", name, err)
			}
		}
	}

	log.Printf("Cleared all cached items")
	return nil
}
