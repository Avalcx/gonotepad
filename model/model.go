package model

import (
	"notepad/config"
	"sync"
	"time"
)

var (
	TextData    = make(map[string]string, 20)
	lastChanged = make(map[string]time.Time, 20)
	mu          sync.RWMutex
)

// 写入后将对应的textName 在这个map中置为true
var CleanList = make(map[string]bool, 20)

// 写入键值
func Set(key, value string) {
	mu.Lock()
	TextData[key] = value
	lastChanged[key] = time.Now()
	mu.Unlock()
}

// 获取键值
func Get(key string) (string, bool) {
	mu.RLock()
	value, exists := TextData[key]
	mu.RUnlock()
	return value, exists
}

// 启动清理器
func StartCleaner() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			cleanExpired()
		}
	}()
}

// 清理过期键
func cleanExpired() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	for key, lastTime := range lastChanged {
		if now.Sub(lastTime) > time.Duration(config.CleanTime)*time.Second {
			delete(TextData, key)
			delete(lastChanged, key)
			println("清理过期键:", key)
		}
	}
}
