package redis_manager

import (
	"sync"
)

var (
	def           Redis
	defaultConfig RedisConfig
	defLocker     sync.Locker
)

func SetDefault(r Redis) {
	def = r
}

// 获取redis连接
func GetConnection() Redis {
	if def != nil {
		return def
	}

	defLocker.Lock()
	defer defLocker.Unlock()

	if def != nil {
		return def
	}

	def = Create(defaultConfig)
	return def
}
