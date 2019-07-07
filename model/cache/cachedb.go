package cache

import (
	. "github.com/fishedee/app/database"
	. "github.com/fishedee/app/ioc"
)

type CacheDb struct {
	db Database
}

func NewCacheDb(db Database) ICacheDb {
	return &CacheDb{
		db: db,
	}
}

func (this *CacheDb) GetByName(name string) []Cache {
	var caches []Cache
	this.db.Where("name = ?", name).MustFind(&caches)
	return caches
}

func (this *CacheDb) Del(cacheId int) {
	this.db.Where("cacheId = ?", cacheId).MustDelete(&Cache{})
}

func (this *CacheDb) Add(cache Cache) {
	this.db.MustInsert(cache)
}

func (this *CacheDb) Mod(cacheId int, cache Cache) {
	this.db.Where("cacheId = ?", cacheId).AllCols().MustUpdate(&cache)
}

func init() {
	MustRegisterIoc(NewCacheDb)
}
