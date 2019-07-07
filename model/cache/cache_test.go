package cache

import (
	. "github.com/fishedee/app/database"
	. "github.com/fishedee/assert"
	"testing"
)

func InitDb(db Database) {
	db.MustExec(`
	create table t_cache(
		cacheId integer primary key autoincrement,
		Name varchar(128) not null,
		value text not null,
		createTime timestamp not null,
		modifyTime timestamp not null
	);
	`)
}

func TestCache(t *testing.T) {
	db := NewDatabaseTest()
	InitDb(db)

	cacheDb := NewCacheDb(db)
	cacheAo := NewCacheAo(cacheDb)

	cacheAo.Set("a", "123")
	AssertEqual(t, cacheAo.Get("a"), "123")

	cacheAo.Set("a", "89")
	AssertEqual(t, cacheAo.Get("a"), "89")

	cacheAo.Del("a")
	AssertEqual(t, cacheAo.Get("a"), "")
}
