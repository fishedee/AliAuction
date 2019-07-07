package cache

import (
	"time"
)

type Cache struct {
	CacheId    int `xorm:"autoincr"`
	Name       string
	Value      string
	CreateTime time.Time `xorm:"created"`
	ModifyTime time.Time `xorm:"updated"`
}
