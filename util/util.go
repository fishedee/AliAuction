package util

import (
	. "aliauction/model/serverchan"
	. "github.com/fishedee/app/config"
	. "github.com/fishedee/app/database"
	. "github.com/fishedee/app/ioc"
	. "github.com/fishedee/app/log"
)

func NewMyConfig() Config {
	config, err := NewConfig("ini", "data/conf/config.ini")
	if err != nil {
		panic(err)
	}
	return config
}

func NewMyLog(config Config) Log {
	var logConfig LogConfig
	config.MustBind("log", &logConfig)
	logger, err := NewLog(logConfig)
	if err != nil {
		panic(err)
	}
	return logger
}

func NewMyDatabase(config Config) Database {
	var databaseConfig DatabaseConfig
	config.MustBind("db", &databaseConfig)
	db, err := NewDatabase(databaseConfig)
	if err != nil {
		panic(err)
	}
	return db
}

func NewMyServerChanConfig(config Config) ServerChanConfig {
	var serverChanConfig ServerChanConfig
	config.MustBind("serverchan", &serverChanConfig)
	return serverChanConfig
}

func init() {
	MustRegisterIoc(NewMyConfig)
	MustRegisterIoc(NewMyLog)
	MustRegisterIoc(NewMyDatabase)
	MustRegisterIoc(NewMyServerChanConfig)
}
