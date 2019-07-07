package cache

import (
	. "github.com/fishedee/app/ioc"
)

type CacheAo struct {
	cacheDb ICacheDb
}

func NewCacheAo(cacheDb ICacheDb) ICacheAo {
	return &CacheAo{
		cacheDb: cacheDb,
	}
}

func (this *CacheAo) Has(key string) bool {
	caches := this.cacheDb.GetByName(key)
	return len(caches) != 0
}

func (this *CacheAo) Get(key string) string {
	caches := this.cacheDb.GetByName(key)
	if len(caches) == 0 {
		return ""
	}
	return caches[0].Value
}

func (this *CacheAo) Del(key string) {
	caches := this.cacheDb.GetByName(key)
	if len(caches) == 0 {
		return
	}
	this.cacheDb.Del(caches[0].CacheId)
}

func (this *CacheAo) Set(key string, value string) {
	caches := this.cacheDb.GetByName(key)
	if len(caches) == 0 {
		this.cacheDb.Add(Cache{
			Name:  key,
			Value: value,
		})
	} else {
		cache := caches[0]
		cache.Value = value
		this.cacheDb.Mod(cache.CacheId, cache)
	}
}

func init() {
	MustRegisterIoc(NewCacheAo)
}
