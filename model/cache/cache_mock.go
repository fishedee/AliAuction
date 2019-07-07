package cache

import ()

type ICacheAo interface {
	Del(key string)
	Get(key string) string
	Has(key string) bool
	Set(key string, value string)
}

type CacheAoMock struct {
	DelHandler func(key string)
	GetHandler func(key string) string
	HasHandler func(key string) bool
	SetHandler func(key string, value string)
}

func (this *CacheAoMock) Del(key string) {
	this.DelHandler(key)
}

func (this *CacheAoMock) Get(key string) string {
	return this.GetHandler(key)
}

func (this *CacheAoMock) Has(key string) bool {
	return this.HasHandler(key)
}

func (this *CacheAoMock) Set(key string, value string) {
	this.SetHandler(key, value)
}

type ICacheDb interface {
	Add(cache Cache)
	Del(cacheId int)
	GetByName(name string) []Cache
	Mod(cacheId int, cache Cache)
}

type CacheDbMock struct {
	AddHandler       func(cache Cache)
	DelHandler       func(cacheId int)
	GetByNameHandler func(name string) []Cache
	ModHandler       func(cacheId int, cache Cache)
}

func (this *CacheDbMock) Add(cache Cache) {
	this.AddHandler(cache)
}

func (this *CacheDbMock) Del(cacheId int) {
	this.DelHandler(cacheId)
}

func (this *CacheDbMock) GetByName(name string) []Cache {
	return this.GetByNameHandler(name)
}

func (this *CacheDbMock) Mod(cacheId int, cache Cache) {
	this.ModHandler(cacheId, cache)
}
